from collections import namedtuple
from struct import Struct
import os

'''
DZH day line data reader
'''


class DzhDayLine:

    def __init__(self, file_path):
        self.file_path = file_path

    def read_data_from_file(self):
        # is file exists?
        if not self.file_path or not os.path.exists(self.file_path):
            return False

        # begin load file
        with open(self.file_path, 'rb') as f:
            pass

        return True


def main():
    # create the reader
    reader = DzhDayLine('c:\\dzh2\\data\\sh\\DAY.DAT')
    # load data
    if not reader.read_data_from_file():
        print('unable read data file file!!!')


if __name__ == '__main__':
    main()


