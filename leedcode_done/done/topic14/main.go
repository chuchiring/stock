package main

import "fmt"

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	minlen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minlen {
			minlen = len(strs[i])
		}
	}

	j := 0
	count := -1
	for j < minlen {
		notSame := false
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] != strs[i][j] {
				notSame = true
				break
			}
		}
		if notSame {
			break
		}
		j++
		count++
	}

	if count == -1 {
		return ""
	}

	return strs[0][0 : count+1]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	minlen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minlen {
			minlen = len(strs[i])
		}
	}

	j := 0
	count := -1
	for j < minlen {
		notSame := false
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] != strs[i][j] {
				notSame = true
				break
			}
		}
		if notSame {
			break
		}
		j++
		count++
	}

	if count == -1 {
		return ""
	}

	return strs[0][0 : count+1]
}

func main() {
	v := []string{"babb", "caa"}
	s := longestCommonPrefix(v)
	fmt.Println(s)
}
