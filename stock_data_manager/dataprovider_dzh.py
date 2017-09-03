from collections import namedtuple, defaultdict
from struct import Struct, calcsize
import os
from datetime import datetime
from dataprovider_def import *

class DzhDataProvider(IStockDataProvider):

    __DAY_FILES__ = {'SH': '\\data\\sh\\DAY.DAT',
                     'SZ': '\\data\\sz\\DAY.DAT'}

    __PWR_FILE__ = '\\download\\PWR\\full.PWR'

    def __init__(self, dzh_install_path: str):
        self.log('load data for DZH v6.0 only!')
        self._stock_headers = {}
        self._stock_xr_rights = {}
        self._dzh_install_path = dzh_install_path

    def prepare(self, stock_manager: IStockManager):

        self._stock_manager = stock_manager

        for market, path in self.__DAY_FILES__.items():
            self.load_stock_headers(market, self._dzh_install_path + path)

        self.log(f'{len(self._stock_headers)} stock header fetched')

        self.load_pwr_from_file(self._dzh_install_path + self.__PWR_FILE__)


    def load_stock_headers(self, market: str, filename: str):

        def read_stock_header(file_stream, market) -> StockHeader:
            # stock code and dayline count
            stock_code = Struct('<10s').unpack_from(file_stream.read(10))[0].decode()[:6]
            stock_dayline_count = Struct('<i').unpack_from(file_stream.read(4))[0]
            # blocks
            blocks = []
            for index in range(25):
                block = Struct('<h').unpack_from(file_stream.read(2))[0]
                if block != -1:
                    blocks.append(block)
            return StockHeader(f'{market}{stock_code}', stock_dayline_count, blocks)

        if not filename or not os.path.exists(filename):
            self.log(f'{filename} not exist!')
            return

        # open file
        with open(filename, 'rb') as f:
            # read header
            buffer = f.read(calcsize('<6i'))
            stock_count = Struct('<6i').unpack_from(buffer)[3]
            self.log(f'Read {filename} , market: {market}, count: {stock_count}')

            # read stock header info
            for index in range(stock_count):
                header = read_stock_header(f, market)
                self._stock_manager.add_stock(header.code)
                self._stock_headers[header.code] = header

    # load the PWR file for "Exclude Right" info
    def load_pwr_from_file(self, file_path):
        self.log('Reading PWR ...')

        with open(file_path, 'rb') as f:

            def read_stock_code():
                buffer = f.read(16)
                if buffer:
                    return Struct('<16s').unpack_from(buffer)[0].decode()[:8]
                else:
                    return None

            def read_rights_info(stock_code: str, xr_list: list):
                buffer = f.read(4)
                if buffer and buffer != b'\xff'*4:
                    xf_info = StockXRInfo(datetime.utcfromtimestamp(Struct('<i').unpack_from(buffer)[0]).date(),
                                          *Struct('<4f').unpack_from(f.read(16)))
                    xr_list.append(xf_info)
                    return True
                else:
                    return False

            # bypass header
            f.seek(12, 0)
            # read stock code
            while True:
                stock_code = read_stock_code()

                if stock_code:
                    xr_list = []
                    while read_rights_info(stock_code, xr_list):
                        pass

                    if xr_list:
                        xr_list.sort()
                        self._stock_xr_rights[stock_code] = xr_list
                else:
                    break

        self.log(f'{len(self._stock_xr_rights)} PWR records loaded')

    def _calc_xr_price(self, xr_rights, daylines):
        # do nothing if have no xr_rights or already calced daylines
        if not xr_rights or not daylines:
            return

        price_factor = 1
        tmp_xr_rights = xr_rights
        cur_xr_right = tmp_xr_rights.pop()

        '''
        LFactorPrice := ((FDayLineEx.OriginClose[AIndex] - LTempRightList.Items[0].Bonus)
          + LTempRightList.Items[0].FitRight * LTempRightList.Items[0].FitPrice)
          /(1 + LTempRightList.Items[0].GiveRight);
        LFactor := LFactor * (LFactorPrice / FDayLineEx.OriginClose[AIndex]);
        LTempRightList.Delete(0);
        '''

        xr_daylines = []

        for dayline in reversed(daylines):
            if cur_xr_right and cur_xr_right.date > dayline.date:
                tmp_factor_price = ((dayline.close - cur_xr_right.bonus) + cur_xr_right.fit_right*cur_xr_right.fit_price) / (1 + cur_xr_right.give_right)
                price_factor = price_factor * (tmp_factor_price / dayline.close)
                #pop next
                cur_xr_right = tmp_xr_rights.pop() if tmp_xr_rights else None

            xr_daylines.append(StockDayLine(dayline.date,
                                                 round(dayline.open * price_factor+0.001,2),
                                                 round(dayline.high * price_factor+0.001,2),
                                                 round(dayline.low * price_factor+0.001,2),
                                                 round(dayline.close * price_factor+0.001,2),
                                                 dayline.vol, dayline.amount))

        xr_daylines.sort()
        return xr_daylines

    def load_stock_dayline(self, stock: IStock, file_stream=None):
        # is file exists?
        file_path = self._dzh_install_path + self.__DAY_FILES__[stock.code[:2]]
        if not file_path or not os.path.exists(file_path):
            return False

        def read_dayline(filestream):
            # first find the stock headers
            header = self._stock_headers.get(stock.code, None)
            if not header:
                return
            # read dayline
            daylines = []
            latest_date = 0
            for block in header.blocks:
                # seek to the block start position
                # every block has 256 day line, every day line has 32 bytes
                filestream.seek(int('0x41000', 16) + 256 * 32 * block, 0)
                # now read the 256 blocks
                for i in range(256):
                    # check the first 4 bytes (date) valid or not
                    buffer = filestream.read(4)
                    if not buffer:
                        break
                    # check is the date less than latest date
                    tmp_date = Struct('i').unpack_from(buffer)[0]
                    if tmp_date <= latest_date:
                        break
                    latest_date = tmp_date
                    # convert unix datetime to python datetime
                    day = datetime.utcfromtimestamp(tmp_date).date()
                    # convert float value
                    day_line = StockDayLine(day, *[x for x in Struct('<6f').unpack_from(filestream.read(24))])
                    # skip the up/down count
                    filestream.read(4)
                    # append dayline to stock object
                    daylines.append(day_line)
            # do XR rights

            xr_rights = self._stock_xr_rights.get(stock.code)
            if xr_rights:
                daylines = self._calc_xr_price(xr_rights, daylines)

            for dayline in daylines:
                stock.add_dayline(dayline)

        if not file_stream:
            handle = open(file_path, 'rb')
            try:
                read_dayline(handle)
            finally:
                handle.close()
        else:
            read_dayline(file_stream)
