from dataprovider_dzh import DzhDataProvider
from stockmanager import StockManager

def test_single_stock():
    stock_manager = StockManager(DzhDataProvider('c:\\dzh2'))

    stock = stock_manager.get_stock('SH600000')

    if not stock:
        return

    stock.export_to_csv('r:\\')

def test_all_stock():
    stock_manager = StockManager(DzhDataProvider('c:\\dzh2'))


    for stock in stock_manager.stocks:
        stock.export_to_csv('r:\\')


if __name__ == '__main__':
    test_single_stock()