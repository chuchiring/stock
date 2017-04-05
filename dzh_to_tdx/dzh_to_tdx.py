#coding=utf-8

import os

def isnumeric(s):
	try:
		float(s)
		return True
	except ValueError:
		pass

def dzh_to_tdx(file):
	tdx_lines = []
	is_number_data = True;
	
	with open(file, 'r') as F:
		for line in F.readlines():
			if line.startswith('SH') or line.startswith('SZ'):
				code, value = line.split("\t", 2)
				if not isnumeric(value):
					is_number_data = False;
					break;

	with open(file, 'r') as F:
		for line in F.readlines():
			if line.startswith('SH') or line.startswith('SZ'):
				code, value = line.split("\t", 2)
				code = '1|' + code[2:] if code.startswith('SH') else '0|' + code[2:]
				if is_number_data:
					tdx_lines.append(f'{code}||{value}')
				else:
					tdx_lines.append(f'{code}|{value}')
				# print(f'{code}||{value}')

	if len(tdx_lines) > 0:
		with open('tdx_' + file, 'w') as F:
			F.writelines(tdx_lines)
				




for file in os.listdir('.'):
	if os.path.isfile(file) and os.path.splitext(file)[1].lower() == '.txt':
		dzh_to_tdx(file)

