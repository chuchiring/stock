from collections import namedtuple
from struct import Struct, calcsize
import os

'''
DZH day line data readers
'''

StockHeader = namedtuple('StockHeader', ['code', 'count', 'blocks'])
StockDayLine = namedtuple('StockDayLine', ['date', 'open', 'high', 'low', 'close', 'vol', 'amount', 'dummy'])


class StockManager:

    def __init__(self):
        self.stocks = {}
        self.stock_headers = []

    @staticmethod
    def read_stock_header(file_handle, market) -> StockHeader:

        # stock code
        stock_code = Struct('<10s').unpack_from(file_handle.read(10))[0].decode().rstrip('\0x00\0x00')
        stock_code = f'{stock_code}.{market}'

        # stock day count
        day_count = Struct('<i').unpack_from(file_handle.read(4))[0]

        # blocks
        blocks = []
        for index in range(25):
            block = Struct('<h').unpack_from(file_handle.read(2))[0]
            if block != -1:
                blocks.append(block)

        return StockHeader(stock_code, day_count, blocks)

    @staticmethod
    def read_stock_day_line(file_handle, stock_header: StockHeader=None):
        # do check
        if not file_handle or not stock_header:
            return

        # read block
        for block in stock_header.blocks:
            # seek to the block start position
            # every block has 256 day line, every day line has 32 bytes
            file_handle.seek(int('0x41000', 16) + 256*32*block, 0)
            day_line = StockDayLine(*Struct('<7fi').unpack_from(file_handle.read(32)))
            print(day_line)

    def read_data_from_file(self, file_path, market):
        # is file exists?
        if not file_path or not os.path.exists(file_path):
            return False

        # begin load file
        with open(file_path, 'rb') as f:

            # read header
            file_header_data = f.read(calcsize('<6i'))
            stock_count = Struct('<6i').unpack_from(file_header_data)[3]
            print(f'{file_path} , stock count is {stock_count}')

            # read stock header info
            for index in range(stock_count):
                stock_header = self.read_stock_header(f, market)
                self.stock_headers.append(stock_header)

            # read stock day line info
            for header in self.stock_headers:
                self.read_stock_day_line(f, header)

        return True


def main():
    # create the reader
    stock_manager = StockManager()

    # load SH data
    if not stock_manager.read_data_from_file('r:\\DAY.DAT', 'SH'):
        print('unable read data file file!!!')


if __name__ == '__main__':
    main()


