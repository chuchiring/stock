package main

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ksimilarity(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("ksimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
