package main

import "fmt"

func shortestSubarray(A []int, K int) int {
	if len(A) == 0 {
		return -1
	}

	//开始正式计算，最后一次全加法和第一次不用计算了
	low, high, arrLen := 0, len(A)-1, len(A)
	lastFoundIndex := arrLen + 1

	for low <= high {
		mid := (low + high) >> 1

		num := 0
		found := false

		//计算连续mid个连续的值
		for i := 0; i < arrLen; i++ {
			if i+mid < arrLen {
				if i == 0 {
					for j := 0; j <= mid; j++ {
						num = num + A[j]
					}
				} else {
					num = num - A[i-1] + A[i+mid]
				}
				//判断总和是不是大于K
				if num >= K {
					found = true
					break
				}
			}
		}

		//继续查找
		if found {
			high = mid - 1
		} else {
			low = mid + 1
		}

		if found && mid < lastFoundIndex {
			lastFoundIndex = mid
		}

	}

	if lastFoundIndex == arrLen+1 {
		return -1
	}

	return lastFoundIndex + 1
}

func main() {
	v := shortestSubarray([]int{44, -25, 75, -50, -38, -42, -32, -6, -40, -47}, 19)

	fmt.Println(v)
}
