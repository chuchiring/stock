from collections import namedtuple
from abc import abstractmethod, ABCMeta

StockHeader = namedtuple('StockHeader', 'code count blocks')

StockDayLine = namedtuple('StockDayLine', 'date open high low close vol amount')

StockXRInfo = namedtuple('StockXRInfo', 'date give_right fit_price fit_right bonus')

class IBaseStockComponent(metaclass=ABCMeta):

    def log(self, log: str):
        print(self.__class__.__name__ + ' :: ' + log)

class IStock(IBaseStockComponent):

    @abstractmethod
    def code(self):
        pass

    @abstractmethod
    def daylines(self) -> []:
        pass

    @abstractmethod
    def add_dayline(self, day_line: StockDayLine):
        pass

    @abstractmethod
    def export_to_csv(self, file_path):
        pass

class IStockManager(IBaseStockComponent):

    @abstractmethod
    def add_stock(self, stock_name: str):
        pass

class IStockDataProvider(IBaseStockComponent):

    @abstractmethod
    def prepare(self, stock_manager: IStockManager):
        pass

    @abstractmethod
    def load_stock_dayline(self, stock: IStock, file_stream=None):
        pass