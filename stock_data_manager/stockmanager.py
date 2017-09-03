from dataprovider_def import IStockDataProvider, IStock, StockDayLine, IStockManager
import csv


class Stock(IStock):

    def __init__(self, code, provider: IStockDataProvider):
        self._daylines = []
        self._code = code
        self._provider = provider
        self._dayline_fetched = False

    @property
    def code(self):
        return self._code

    @property
    def daylines(self) -> []:
        if not self._dayline_fetched and not self._daylines:
            self._dayline_fetched = True
            self._provider.load_stock_dayline(self)

        return self._daylines

    def add_dayline(self, day_line: StockDayLine):
        self._daylines.append(day_line)

    def export_to_csv(self, file_path):
        self.log(f'writing {file_path}{self.code}.csv...')
        with open(f'{file_path}{self.code}.csv', 'w') as f:
            writer = csv.writer(f)
            writer.writerow(('date', 'open', 'high', 'low', 'close', 'vol', 'amount'))
            for line in self.daylines:
                writer.writerow((line))

class StockManager(IStockManager):
    # init
    def __init__(self, provider: IStockDataProvider):
        self._provider = provider
        # _stocks item contains (stockid, Stock)
        self._stocks = {}
        # load stock basic info
        self.init_provider()


    def add_stock(self, stock_code: str):
        self._stocks[stock_code] = Stock(stock_code, self._provider)

    @property
    def stocks(self) -> []:
        return list(self._stocks.values())

    def get_stock(self, stock_code) -> Stock:
        stock = self._stocks.get(stock_code.upper())
        if not stock:
            self.log(f'invalid stock id: {stock_code}')
        return stock

    # fill stock data
    def init_provider(self):
        self.log('prepare data...')
        self._provider.prepare(stock_manager=self)

