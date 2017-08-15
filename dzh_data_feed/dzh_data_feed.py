from collections import namedtuple
from struct import Struct, calcsize
import os
from datetime import datetime
import csv

'''
DZH day line data readers
'''

StockHeader = namedtuple('StockHeader', ['code', 'count', 'blocks'])
StockDayLine = namedtuple('StockDayLine', ['date', 'open', 'high', 'low', 'close', 'vol', 'amount'])

class Stock:

    def __init__(self, code: str):
        self.code = code
        self.origin_daylines = []

    def add_one_dayline(self, dayline: StockDayLine):
        self.origin_daylines.append(dayline)

    def export_as_csv(self, file_path):
        print(f'writing {file_path}{self.code}.csv...')
        with open(f'{file_path}{self.code}.csv', 'w') as f:
            writer = csv.writer(f)
            writer.writerow(('date', 'open', 'high', 'low', 'close', 'vol', 'amount'))
            for line in self.origin_daylines:
                writer.writerow((line))


class StockManager:

    def __init__(self, test_code: str=None):
        self.stocks = {}
        self.stock_headers = []
        self.test_code = test_code

    @staticmethod
    def read_stock_header(file_stream, market) -> StockHeader:

        # stock code
        stock_code = Struct('<10s').unpack_from(file_stream.read(10))[0].decode()[:6]
        stock_code = f'{stock_code}.{market}'

        # stock day count
        day_count = Struct('<i').unpack_from(file_stream.read(4))[0]

        # blocks
        blocks = []
        for index in range(25):
            block = Struct('<h').unpack_from(file_stream.read(2))[0]
            if block != -1:
                blocks.append(block)

        return StockHeader(stock_code, day_count, blocks)

    def read_stock_day_line(self, file_stream, stock_header: StockHeader=None):
        # do check
        if not file_stream or not stock_header:
            return

        # for test
        if self.test_code and stock_header.code != self.test_code:
            return

        print(f'process {stock_header.code}....')

        # create stock object
        stock = Stock(stock_header.code)

        # read block
        latest_date = 0

        for block in stock_header.blocks:
            # seek to the block start position
            # every block has 256 day line, every day line has 32 bytes
            file_stream.seek(int('0x41000', 16) + 256*32*block, 0)

            # now read the 256 blocks
            for i in range(256):
                # check the first 4 bytes (date) valid or not
                raw_data = file_stream.read(4)
                if not raw_data:
                    break
                # check is the date less than latest date
                tmp_date = Struct('i').unpack_from(raw_data)[0]
                if tmp_date <= latest_date:
                    break
                latest_date = tmp_date
                # convert unix datetime to python datetime
                day = datetime.fromtimestamp(tmp_date)

                # convert float value
                day_line = StockDayLine(day, *[round(x+0.001,2) for x in Struct('<6f').unpack_from(file_stream.read(24))])

                # skip the up/down count
                file_stream.read(4)

                # append dayline to stock object
                stock.add_one_dayline(day_line)


        self.stocks[stock_header.code] = stock

    def read_data_from_file(self, file_path, market):
        # is file exists?
        if not file_path or not os.path.exists(file_path):
            return False

        # begin load file
        with open(file_path, 'rb') as f:

            # read header
            file_header_data = f.read(calcsize('<6i'))
            stock_count = Struct('<6i').unpack_from(file_header_data)[3]
            print(f'{file_path} , stock count: {stock_count}')

            # read stock header info
            for index in range(stock_count):
                stock_header = self.read_stock_header(f, market)
                self.stock_headers.append(stock_header)

            # read stock day line info
            for header in self.stock_headers:
                self.read_stock_day_line(f, header)

        return True


def main():

    test_code = None

    # create the reader
    stock_manager = StockManager(test_code)

    # load SH data
    if not stock_manager.read_data_from_file('d:\\DAY.DAT', 'SH'):
        print('unable read data file file!!!')

    # save csv
    if test_code:
        stock_manager.stocks[test_code].export_as_csv('r:\\')
    else:
        for _, stock in stock_manager.stocks.items():
            stock.export_as_csv('r:\\dcu\\')

if __name__ == '__main__':
    main()


