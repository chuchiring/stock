package main

import "fmt"

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	arr := make([]int, len(s))

	//first change '(' to 1 and change ')' to -1
	for i := range arr {
		switch s[i] {
		case '(':
			arr[i] = -1
		default:
			arr[i] = 1
		}
	}

	//now scan array to find 1,-1 pair and set to zero
	for i := range arr {
		if i < len(arr)-1 && arr[i] == -1 && arr[i+1] == 1 {
			arr[i], arr[i+1] = 0, 0
		}
	}

	//define
	fScan := func(i, step int, arr []int) (int, bool) {
		k := i + step
		for k >= 0 && k < len(arr) {
			if arr[k] != 0 {
				if arr[k] == step {
					return k, true
				}
				return k, false
			}
			k = k + step
		}
		return k, false
	}

	i := 0
	for i < len(arr) {
		if arr[i] == 0 {
			li, ri := i, i

			for {
				li, lok := fScan(li, -1, arr)
				ri, rok := fScan(ri, 1, arr)

				if lok && rok {
					arr[li], arr[ri] = 0, 0
					continue
				} else {
					i = ri
					break
				}
			}
		} else {
			i++
		}
	}

	var count, maxcount int
	for i := range arr {
		if arr[i] == 0 {
			count++
			if count > maxcount {
				maxcount = count
			}
		} else {
			count = 0
		}
	}

	// fmt.Println(arr)

	return maxcount
}

func main() {
	// s := "((()())))(()(((())"
	s := "(()"

	fmt.Println(longestValidParentheses(s))
}
