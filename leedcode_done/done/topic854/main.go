package main

import (
	"fmt"
)

func ksimilarity(A string, B string) int {
	debugLog := true

	dst := make([]byte, len(A))
	src := make([]byte, len(dst))

	//remove dupe chars
	arrLen := 0
	for i := range A {
		if A[i] != B[i] {
			dst[arrLen], src[arrLen] = byte(A[i])-96, byte(B[i])-96
			arrLen++
		}
	}

	//define a function to print the packed result
	fPackArray := func(arr []byte) []byte {
		arr1, j := make([]byte, len(arr)), 0
		for i := range arr {
			if arr[i] != 0 {
				arr1[j] = arr[i]
				j++
			}
		}
		return arr1
	}

	//define a log func to print the result of array
	fLogInfo := func() {
		if debugLog {
			fmt.Println(fPackArray(src))
			fmt.Println(fPackArray(dst))
			fmt.Println("----------------------------------------")
		}
	}

	//the min count
	count := 0

	//clean {1,4}{4,1}, swap 1 and clean 2
	fCleanTwoLogic := func() {
		for i := 0; i < arrLen-1; i++ {
			if src[i] != 0 {
				vsrc, vdst := src[i], dst[i]
				for j := i + 1; j < arrLen; j++ {
					if src[j] > 0 && src[j] == vdst && dst[j] == vsrc {
						src[i], dst[i], src[j], dst[j] = 0, 0, 0, 0
						fLogInfo()
						count++
						break
					}
				}
			}
		}
	}

	//clean {1,3}{2,1}{3,1}, swap 2 and clean 3
	fCleanThreeMode := func() {
		for i := 0; i < arrLen-1; i++ {
			if src[i] > 0 {
				secondindex, threeindex := -1, -1
				var vsrc, vdst byte
				for j := 0; j < arrLen; j++ {
					if j != i {
						if src[j] > 0 && (secondindex == -1) && (src[i] == dst[j] || dst[i] == src[j]) {
							if src[i] == dst[j] {
								vsrc, vdst = dst[j], src[i]
							} else {
								vsrc, vdst = src[i], dst[j]
							}
							secondindex = j
						} else if src[j] > 0 && secondindex > 0 && vsrc == dst[j] && vdst == src[j] {
							threeindex = j
							break
						}
					}
				}
				if secondindex > 0 && threeindex > 0 {
					src[i], dst[i] = 0, 0
					src[secondindex], dst[secondindex] = 0, 0
					src[threeindex], dst[threeindex] = 0, 0
					fLogInfo()
					count = count + 2
				}
			}
		}
	}
	fCleanTwoLogic()
	fCleanThreeMode()

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
				fLogInfo()

				count++
			}
		}
		// fCleanTwoLogic()
		// fCleanThreeMode()
		i++
	}

	return count
}

func main() {
	fmt.Println(ksimilarity("abcdefabcdefabcdef", "faebbcddaceeffcbda"))
}
