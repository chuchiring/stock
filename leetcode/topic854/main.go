package main

import (
	"fmt"
)

func ksimilarity(A string, B string) int {
	dst := make([]byte, len(A))
	src := make([]byte, len(dst))

	arrLen := 0
	for i := range A {
		if A[i] != B[i] {
			dst[arrLen], src[arrLen] = byte(A[i]), byte(B[i])
			arrLen++
		}
	}

	var i, count int
	for i < arrLen {
		if src[i] != dst[i] {
			for j := i + 1; j < len(dst); j++ {
				if dst[j] == src[i] {
					dst[i], dst[j] = dst[j], dst[i]
					count++
					break
				}
			}
		}
		i++
	}

	return count
}

func main() {
	// fmt.Println(ksimilarity("bccaba", "abacbc"))
	fmt.Println(ksimilarity("bccaa", "abacc"))
	// fmt.Println(ksimilarity("aabc", "abca"))
}
