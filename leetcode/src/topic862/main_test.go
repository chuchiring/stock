package main

import "testing"

func Test_shortestSubarray(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"s1", args{[]int{1}, 1}, 1},
		{"s2", args{[]int{1, 2}, 4}, -1},
		{"s3", args{[]int{2, -1, 2}, 3}, 3},
		{"s4", args{[]int{17, 85, 93, -45, -21}, 150}, 2},
		{"s5", args{[]int{56, -21, 56, 35, -9}, 61}, 2},
		{"s6", args{[]int{44, -25, 75, -50, -38, -42, -32, -6, -40, -47}, 19}, 1},
		{"s7", args{[]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453}, 9},
		{"s8", args{[]int{1, -1, -1, -1, -1}, 1}, 1},
		{"k8", args{[]int{58, 6, -14, 3, -14, 3, -11, 13, 22}, 55}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSubarray(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("shortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
