package main

import (
	"fmt"
	"sort"
)

type (
	//Point include x, y
	Point struct {
		X, Y int
	}
	pointsSorter struct {
		data []Point
	}
)

func (p *pointsSorter) Len() int      { return len(p.data) }
func (p *pointsSorter) Swap(i, j int) { p.data[i], p.data[j] = p.data[j], p.data[i] }
func (p *pointsSorter) Less(i, j int) bool {
	if p.data[i].X == p.data[j].X {
		return p.data[i].Y < p.data[j].Y
	}
	return p.data[i].X < p.data[j].X
}

func maxInt(first int, param ...int) int {
	if len(param) == 0 {
		return first
	}

	for _, v := range param {
		if v > first {
			first = v
		}
	}
	return first
}

func max2Int(first, second int, param ...int) (int, int) {
	if len(param) == 0 {
		return first, second
	}

	for _, v := range param {
		if v > first || v > second {
			if first > second {
				second = v
			} else {
				first = v
			}
		}
	}
	return first, second
}

func maxPoints(points []Point) int {
	// debuglog := true

	if len(points) <= 2 {
		return len(points)
	}

	maxLen := 2

	arr := &pointsSorter{points}
	sort.Sort(arr)

	//分离x，y，用于统计x,y总相同数子的最大值，这是十字直线
	xarr, yarr := make([]int, len(points)), make([]int, len(points))
	for i := range points {
		xarr[i], yarr[i] = points[i].X, points[i].Y
	}

	fCalcFlatLen := func(arr []int) int {
		sort.Ints(arr)
		max, cur, num := 1, 1, arr[0]
		for i := 1; i < len(arr); i++ {
			if num == arr[i] {
				cur++
				if i == len(arr)-1 {
					max = maxInt(cur, max)
				}
			} else {
				num = arr[i]
				max = maxInt(cur, max)
				cur = 1
			}
		}
		return max
	}

	//计算正斜线和反斜线
	fXX := func(arr []Point, mode int) int {
		max := 2
		for i := range arr {
			srcP, nlen := arr[i], 1
			for j := i + 1; j < len(arr); j++ {
				switch mode {
				case 1:
					if srcP.X-arr[j].X == srcP.Y-arr[j].Y {
						nlen++
					}
				case -1:
					if srcP.X-arr[j].X == arr[j].Y-srcP.Y {
						nlen++
					}
				}
			}
			max = maxInt(max, nlen)
		}
		return max
	}

	//计算相同点
	fSamePoints := func(arr []Point) int {
		dupe1, dupe2, cur, p := 1, 1, 1, arr[0]
		for i := 1; i < len(arr); i++ {
			if p.X == arr[i].X && p.Y == arr[i].Y {
				cur++
				if i == len(arr)-1 {
					dupe1, dupe2 = max2Int(dupe1, dupe2, cur)
				}
			} else {
				p = arr[i]
				dupe1, dupe2 = max2Int(dupe1, dupe2, cur)
				cur = 1
			}
		}
		return dupe1 + dupe2
	}

	return maxInt(maxLen, fCalcFlatLen(xarr), fCalcFlatLen(yarr),
		fXX(points, 1), fXX(points, -1), fSamePoints(points))
}

func main() {
	// points := []point{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	points := []Point{{1, 1}, {1, 1}, {1, 1}}
	fmt.Println("p:", points)
	fmt.Println("直线最大: ", maxPoints(points))
}
