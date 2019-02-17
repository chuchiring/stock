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

func maxPoints(points []Point) int {
	if len(points) <= 1 {
		return len(points)
	}

	debuglog := true
	arr := &pointsSorter{points}
	sort.Sort(arr)

	xarr, yarr := make([]int, len(points)), make([]int, len(points))
	for i := range points {
		xarr[i], yarr[i] = points[i].X, points[i].Y
	}

	fCalcMaxLenOfXY := func(arr []int) int {
		if len(arr) <= 1 {
			return len(arr)
		}
		sort.Ints(arr)
		maxCount, curCount, num := 1, 1, arr[0]
		for i := 1; i < len(arr); i++ {
			if num != arr[i] {
				if curCount > maxCount {
					maxCount = curCount
				}
				num, curCount = arr[i], 0
				continue
			}
			curCount++
			if curCount > maxCount {
				maxCount = curCount
			}
		}
		return maxCount
	}

	rowMax, colMax := fCalcMaxLenOfXY(xarr), fCalcMaxLenOfXY(yarr)

	//计算斜线
	fXX := func(arr []Point, mode int) int {
		maxcount := 2
		for i := range arr {
			pstart := arr[i]
			curLen := 1
			for j := i + 1; j < len(arr); j++ {
				//判断是否在直线
				if (arr[j].X-pstart.X == arr[j].Y-pstart.Y) || (arr[j].X == pstart.X && arr[j].Y == pstart.Y) {
					curLen++
				}
			}
			if curLen > maxcount {
				maxcount = curLen
			}
		}
		return maxcount
	}

	zxMax := fXX(points, 1)

	fFX := func(arr []Point, mode int) int {
		maxcount := 2
		for i := len(arr) - 1; i >= 0; i-- {
			pstart := arr[i]
			curLen := 1
			for j := i - 1; j >= 0; j-- {
				//判断是否在直线
				if (arr[j].X-pstart.X == pstart.Y-arr[j].Y) || (arr[j].X == pstart.X && arr[j].Y == pstart.Y) {
					curLen++
				}
			}
			if curLen > maxcount {
				maxcount = curLen
			}
		}
		return maxcount
	}

	fxMax := fFX(points, 1)

	if debuglog {
		fmt.Println("a:", *arr)
		fmt.Println("y:", yarr)
		fmt.Println("x:", xarr)
		fmt.Printf("列直线: %d, 行直线: %d, 正斜: %d, 反斜: %d\n", rowMax, colMax, zxMax, fxMax)
	}

	max := colMax
	if rowMax > max {
		max = rowMax
	}

	if zxMax > max {
		max = zxMax
	}

	if fxMax > max {
		max = fxMax
	}
	return max
}

func main() {
	// points := []point{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	points := []Point{{1, 1}, {1, 1}, {2, 3}}
	fmt.Println("p:", points)
	fmt.Println("直线最大: ", maxPoints(points))
}
