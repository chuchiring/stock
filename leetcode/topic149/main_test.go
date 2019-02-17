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
		{"simple", args{[]Point{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
