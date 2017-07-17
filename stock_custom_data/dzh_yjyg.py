import xlrd
from collections import namedtuple

RowData = namedtuple('RowData',
                     ['code', 'name', 'detail', 'info', 'date', 'dummy', 'kb_date', 'kb_last_year', 'kb_this_year',
                      'standard_date', 'standard_increase', 'standard_money'])


def money_to_str(total):
    if abs(total) >= 10000 * 10000:
        return '%.2f亿' % (total / 10000 / 10000)

    elif abs(total) >= 10000:
        return '%d万' % (total / 10000)
    else:
        return '%d元' % total


def main():
    # 数据文件名称
    data_file = 'r:\\业绩预告.xls'
    dzh_file = 'r:\\业绩预告.txt'
    txt_lines = []

    # 读取xls
    book = xlrd.open_workbook(data_file)

    def format_cell_datetime(cell_date_value: float):
        date_tuple = xlrd.xldate_as_tuple(cell_date_value, book.datemode)
        return f'{date_tuple[0]}-{date_tuple[1]}-{date_tuple[2]}'

    # 定位sheet
    sheet = book.sheet_by_index(0)

    current_row = 1
    num_rows = sheet.nrows

    while current_row < num_rows:
        row = sheet.row(current_row)

        # 是股票代码
        if sheet.cell(current_row, 0).value.startswith(('0', '6')):
            stock_obj = RowData(*tuple(row))

            # 构建大智慧代码
            stock_code = stock_obj.code.value[:6]
            dzh_code = f'SH{stock_code}' if stock_code.startswith('6') else f'SZ{stock_code}'

            # detail
            detail = ''

            if stock_obj.standard_date.value:
                detail = '%s\t%s 中报业绩 %2.f%%, 净利润 %s\n' % (
                dzh_code, format_cell_datetime(stock_obj.standard_date.value), stock_obj.standard_increase.value,
                money_to_str(stock_obj.standard_money.value))
                print(detail)

            elif stock_obj.kb_date.value and stock_obj.kb_this_year.value:

                if stock_obj.kb_last_year.value:
                    detail = '%s\t%s 中报业绩快报 净利润增长: %2.f%%, %s元\n' % (dzh_code,
                                                               format_cell_datetime(stock_obj.kb_date.value),
                                                               (stock_obj.kb_this_year.value - stock_obj.kb_last_year.value) * 100 / stock_obj.kb_last_year.value,
                                                                money_to_str(stock_obj.kb_this_year.value))
                else:
                    detail = '%s\t%s 中报业绩快报 净利润 %s元\n' % (dzh_code,
                                                               format_cell_datetime(stock_obj.kb_date.value),
                                                                money_to_str(stock_obj.kb_this_year.value))
                print(detail)

            elif stock_obj.date.value:
                detail = '%s\t%s 中报业绩%s %s\n' % (dzh_code, format_cell_datetime(stock_obj.date.value), stock_obj.info.value, stock_obj.detail.value)
                print(detail)

            if len(detail) > 0:
                txt_lines.append(detail)

        current_row += 1

    # 保存为文本
    with open(dzh_file, 'w') as f:
        f.writelines(txt_lines)


if __name__ == '__main__':
    main()