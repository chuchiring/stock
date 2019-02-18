package main

import "fmt"

func shortestSubarray(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}

	minLen := -1

	resultArr := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		resultArr[i] = A[i]
		if resultArr[i] >= K {
			return 1
		}
	}

	deepth := 1
	for deepth < len(A) {
		for i := 0; i < len(A); i++ {
			if i+deepth < len(A) {
				resultArr[i] = resultArr[i] + A[i+deepth]
				if resultArr[i] >= K {
					return deepth + 1
				}
			}
		}
		deepth++
	}

	return minLen
}

func main() {
	v := shortestSubarray([]int{2, -1, 2}, 3)

	fmt.Println(v)
}
