package main

import (
	"testing"
)

func TestTag(t *testing.T) {

	type testInfo = struct {
		tag    string
		result int
	}

	var tags = []testInfo{
		{"(((", 0},
		{"(()", 2},
		{"(())", 4},
		{")()())", 4},
		{")()()))()())", 4},
		{")()()))()()())", 6},
		{"()(())", 6},
		{"(()()))", 6},
	}

	for index, v := range tags {
		i := longestValidParentheses(v.tag)
		if i != v.result {
			t.Fatalf("%d: %s expect %d, but result is %d", index, v.tag, v.result, i)
		}
	}
}
