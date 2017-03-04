#coding=utf-8

import os

def dzh_to_tdx(file):
	tdx_lines = []

	with open(file, 'r') as F:
		for line in F.readlines():
			if line.startswith('SH') or line.startswith('SZ'):
				code, value = line.split("\t", 2)
				code = '1|' + code[2:] if code.startswith('SH') else '0|' + code[2:]
				tdx_lines.append(f'{code}||{value}')
				# print(f'{code}||{value}')

	if len(tdx_lines) > 0:
		with open('tdx_' + file, 'w') as F:
			F.writelines(tdx_lines)
				




for file in os.listdir('.'):
	if os.path.isfile(file) and os.path.splitext(file)[1].lower() == '.txt':
		dzh_to_tdx(file)

