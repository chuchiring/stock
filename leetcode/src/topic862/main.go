package main

import "fmt"

func shortestSubarraySlow(A []int, K int) int {
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

func shortestSubarray(A []int, K int) int {
	if len(A) == 0 {
		return -1
	}

	lastFoundIndex := -1
	low, high, arrLen := 0, len(A)-1, len(A)

	for low <= high {
		mid := (low + high) >> 1
		found := false

		num := 0
		fakeK := K
		for i := 0; i+mid < arrLen; i++ {
			if i == 0 {
				for j := i; j <= i+mid; j++ {
					num = num + A[j]
					if A[j] < 0 {
						fakeK = fakeK + A[j]
					}
				}
			} else {
				num = num - A[i-1] + A[i+mid]
				if A[i-1] < 0 {
					fakeK = fakeK - A[i-1]
				}
				if A[i+mid] < 0 {
					fakeK = fakeK + A[i+mid]
				}
			}

			if num >= fakeK {
				found = true
				lastFoundIndex = mid + 1
				break
			}
		}

		if found {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return lastFoundIndex
}

func main() {
	v := shortestSubarray([]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453)

	fmt.Println(v)
}
