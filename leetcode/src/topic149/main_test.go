package main

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple1", args{[]Point{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}}, 4},
		{"simple2", args{[]Point{{2, 3}, {3, 3}, {-5, 3}}}, 3},
		{"simple3", args{[]Point{{1, 1}, {1, 1}, {2, 3}}}, 3},
		{"simple4", args{[]Point{{3, 10}, {0, 2}, {0, 2}, {3, 10}}}, 4},
		{"simple3", args{[]Point{{1, 1}, {1, 1}, {1, 1}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
