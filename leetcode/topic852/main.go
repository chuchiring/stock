package main

import (
	"fmt"
	"sort"
)

type intSlice2 struct {
	data      []int
	swapcount int
}

func (p *intSlice2) Len() int           { return len(p.data) }
func (p *intSlice2) Less(i, j int) bool { return p.data[i] < p.data[j] }
func (p *intSlice2) Swap(i, j int) {
	p.data[i], p.data[j], p.swapcount = p.data[j], p.data[i], p.swapcount+1
}

func ksimilarity(A string, B string) int {
	arr1 := &intSlice2{data: make([]int, len(A))}
	arr2 := &intSlice2{data: make([]int, len(B))}

	fConvertToArray := func(arr []int, str string) {
		for i := range str {
			arr[i] = int(str[i])
		}
	}

	fConvertToArray(arr1.data, A)
	fConvertToArray(arr2.data, B)

	sort.Sort(arr1)
	sort.Sort(arr2)

	if arr2.swapcount > arr1.swapcount {
		return arr2.swapcount - arr1.swapcount
	}
	return arr1.swapcount - arr2.swapcount
}

func main() {
	fmt.Println(ksimilarity("bccaba", "abacbc"))
	// fmt.Println(ksimilarity("aabc", "abca"))
}
