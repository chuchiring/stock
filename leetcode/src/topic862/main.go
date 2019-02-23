package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func doRandomTest(count int) {
	j := 0
	for j <= count {
		rand.Seed(time.Now().UTC().UnixNano())
		arrlen := rand.Intn(20) + 1
		arr := make([]int, arrlen)
		for i := 0; i < arrlen; i++ {
			arr[i] = rand.Intn(60) - rand.Intn(60) + rand.Intn(60) - rand.Intn(60)
		}

		k := rand.Intn(100)
		fmt.Println(arr, k)

		v := shortestSubarray(arr, k)

		if vok := shortestSubarraySlow(arr, k); vok != v {
			fmt.Printf("Error, should be %d, but got %d\n", vok, v)

			s := ""
			for _, v := range arr {
				s = s + fmt.Sprintf(", %d", v)
			}
			if len(s) > 0 {
				s = s[2:]
			}
			s = fmt.Sprintf("[]int{%s}, %d", s, k)
			fmt.Printf("shortestSubarray(%s)\n", s)
			fmt.Printf("{\"k8\", args{%s}, %d},\n", s, vok)
			//{"s8", args{[]int{1, -1, -1, -1, -1}, 1}, 1},
			log.Fatal("")
		} else {
			fmt.Printf("Success, result is %d\n", v)
		}

		j++

		time.Sleep(time.Millisecond * 50)
	}
}

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

	lastIndex, lastValue := -1, -1

	for low <= high {
		mid := (low + high) >> 1
		found := false

		num := 0
		max := 0
		for i := 0; i+mid < arrLen; i++ {
			if i == 0 {
				for j := i; j <= i+mid; j++ {
					num = num + A[j]
				}
				max = num
			} else {
				num = num - A[i-1] + A[i+mid]
				if num > max {
					max = num
				}
			}

			if num >= K {
				found = true
				lastFoundIndex = mid + 1
				break
			}
		}

		if found {
			high = mid - 1
			lastIndex, lastValue = -1, -1
		} else {
			if lastIndex == -1 {
				lastIndex, lastValue = mid, max
			} else {
				if max <= lastValue && mid > lastIndex {
					high = lastIndex
					low = 0
					continue
				}

				lastIndex, lastValue = mid, max
			}

			low = mid + 1
		}
	}

	return lastFoundIndex
}

func main() {
	// v := shortestSubarray([]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453)
	// fmt.Println(v)
	// doRandomTest(1000)
	v := shortestSubarray([]int{1, -1, -1, -1, -1}, 1)
	fmt.Println(v)
}
