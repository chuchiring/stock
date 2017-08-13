from collections import namedtuple
from struct import Struct, calcsize
import os

'''
DZH day line data readers
'''


class DzhDayLine:

    def __init__(self, file_path):
        self.file_path = file_path
        self.stock_count = 0
        self.stocks = {}

    def read_data_from_file(self):
        # is file exists?
        if not self.file_path or not os.path.exists(self.file_path):
            return False

        # begin load file
        with open(self.file_path, 'rb') as f:

            # read header
            file_header_data = f.read(calcsize('<6i'))
            file_header = Struct('<6i').unpack_from(file_header_data)
            self.stock_count = file_header[3]
            print(f'{self.file_path} , stock count is {self.stock_count}')

            # read stock header info
            for index in range(self.stock_count):
                stock_header_data = f.read(calcsize('<10sI25H'))
                stock_header = Struct('<10sI25H').unpack_from(stock_header_data)
                stock_code = stock_header[0].decode()
                print(stock_code)


        return True


def main():
    # create the reader
    reader = DzhDayLine('r:\\DAY.DAT')

    # load data
    if not reader.read_data_from_file():
        print('unable read data file file!!!')


if __name__ == '__main__':
    main()


