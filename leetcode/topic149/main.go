package main

import (
	"fmt"
	"sort"
)

type (
	//Point include x, y
	Point struct {
		x, y int
	}
	pointsSorter struct {
		data []Point
	}
)

func (p *pointsSorter) Len() int      { return len(p.data) }
func (p *pointsSorter) Swap(i, j int) { p.data[i], p.data[j] = p.data[j], p.data[i] }
func (p *pointsSorter) Less(i, j int) bool {
	if p.data[i].x == p.data[j].x {
		return p.data[i].y < p.data[j].y
	}
	return p.data[i].x < p.data[j].x
}

func maxPoints(points []Point) int {
	debuglog := true
	arr := &pointsSorter{points}
	sort.Sort(arr)

	xarr, yarr := make([]int, len(points)), make([]int, len(points))
	for i := range points {
		xarr[i], yarr[i] = points[i].x, points[i].y
	}

	fCalcMaxLenOfXY := func(arr []int) int {
		sort.Ints(arr)
		maxCount, curCount, num := 0, 0, -1
		for _, v := range arr {
			if num != v {
				if curCount > maxCount {
					maxCount = curCount
				}
				num, curCount = v, 0
				continue
			}
			curCount++
		}
		return maxCount + 1
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
				if arr[j].x-pstart.x == arr[j].y-pstart.y {
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
				if arr[j].x-pstart.x == pstart.y-arr[j].y {
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
	points := []Point{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println("p:", points)
	fmt.Println("直线最大: ", maxPoints(points))
}
