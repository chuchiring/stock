package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastFoundIndex := 0
	low, high, arrLen := 0, len(nums)-1, len(nums)

	for low <= high {
		mid := (low + high) >> 1
		found := false

		num := 0
		for i := 0; i+mid < arrLen; i++ {
			if i == 0 {
				for j := i; j <= i+mid; j++ {
					num = num + nums[j]
				}

			} else {
				num = num - nums[i-1] + nums[i+mid]
			}

			if num >= s {
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
	v := minSubArrayLen(700, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(v)
}
