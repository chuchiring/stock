package main

import "testing"

/*
解题思路

1
找出两个数组中交叉相等的进行交换， 这样一次可以消掉2个

2
找出两个数组中2次移动可以消掉3个的

3
最后正常单次每次消掉1个
*/

func Test_ksimilarity(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"simple", args{"ab", "ba"}, 1},
		{"simple2", args{"aabc", "abca"}, 2},
		{"simple2", args{"bccaba", "abacbc"}, 3},
		{"simple2", args{"abccaacceecdeea", "bcaacceeccdeaae"}, 9},
		{"simple2", args{"aabbccddee", "cbeddebaac"}, 6},
		{"simple2", args{"cdacdfdcacbaedfdaeaf", "afdcacedcdabdeffcaad"}, 10},
		{"simple2", args{"aaaabbbbccccddddeeee", "aceecdabdcdbebeaadbc"}, 10},
		{"simple2", args{"abcdefabcdefabcdef", "faebbcddaceeffcbda"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ksimilarity(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("ksimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
