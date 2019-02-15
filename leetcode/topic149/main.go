package main

import (
	"fmt"
	"sort"
)

type (
	point struct {
		x, y int
	}
	pointsSorter struct {
		data []point
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

func maxPoints(points []point) int {
	// arr := &pointsSorter(points)
	// sort.Sort(arr)

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

	fmt.Println("y:", yarr)
	fmt.Println("x:", xarr)
	fmt.Printf("列直线: %d, 行直线: %d\n", rowMax, colMax)

	if colMax > rowMax {
		return colMax
	}
	return colMax
}

func main() {
	points := []point{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	fmt.Println("p:", points)
	fmt.Println("直线最大: ", maxPoints(points))
}
