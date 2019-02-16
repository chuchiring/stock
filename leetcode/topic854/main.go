package main

import (
	"fmt"
)

func ksimilarity(A string, B string) int {
	debugLog := true

	dst := make([]byte, len(A))
	src := make([]byte, len(dst))

	arrLen := 0
	for i := range A {
		if A[i] != B[i] {
			dst[arrLen], src[arrLen] = byte(A[i])-96, byte(B[i])-96
			arrLen++
		}
	}

	if debugLog {
		fmt.Println(src)
		fmt.Println(dst)
		fmt.Println("======================================")
	}

	count := 0

	//清除上下一样的
	fCleanSameLogic := func() {
		for i := 0; i < arrLen-1; i++ {
			vsrc, vdst := src[i], dst[i]
			for j := i + 1; j < arrLen; j++ {
				if src[j] > 0 && src[j] == vdst && dst[j] == vsrc {
					src[i], dst[i], src[j], dst[j] = 0, 0, 0, 0
					if debugLog {
						fmt.Println(src)
						fmt.Println(dst)
						fmt.Println("----------------")
					}
					count++
					break
				}
			}
		}
	}

	fCleanSameLogic()

	if debugLog {
		fmt.Println("\n enter single mode")
		fmt.Println("")
	}

	var i int
	for i < arrLen {
		if src[i] > 0 && src[i] != dst[i] {
			match1 := 0
			for j := i + 1; j < arrLen; j++ {
				if dst[j] == src[i] {
					match1 = j
					if dst[i] == src[j] {
						break
					}
				}

			}
			j := match1
			if dst[j] == src[i] {
				dst[i], dst[j] = dst[j], dst[i]

				//判断交换后的位置是否相等
				if dst[j] == src[j] {
					dst[j], src[j] = 0, 0
				}

				src[i], dst[i] = 0, 0
				if debugLog {
					fmt.Println(src)
					fmt.Println(dst)
					fmt.Println("----------------")
				}

				count++
			}
		}
		fCleanSameLogic()
		i++
	}

	return count
}

func main() {
	fmt.Println(ksimilarity("aaaabbbbccccddddeeee", "aceecdabdcdbebeaadbc"))
}
