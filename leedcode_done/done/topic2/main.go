package main

import "fmt"

//ListNode = 12
type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	root := &ListNode{Val: nums[0]}
	node := root
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{Val: nums[i]}
		node = node.Next
	}
	return root
}

func printList(node *ListNode) []int {
	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	return nums
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	fCalcLen := func(root *ListNode) int {
		count := 0
		for root != nil {
			root = root.Next
			count++
		}
		return count
	}

	var long, short, longroot *ListNode
	if fCalcLen(l1) > fCalcLen(l2) {
		long, short, longroot = l1, l2, l1
	} else {
		long, short, longroot = l2, l1, l2
	}

	vup := 0
	for short != nil {
		v := long.Val + short.Val + vup
		vup = v / 10
		long.Val = v % 10

		short = short.Next
		if long.Next == nil && vup > 0 {
			long.Next = &ListNode{}
		}
		long = long.Next
	}

	for vup > 0 {
		v := long.Val + vup
		vup = v / 10
		long.Val = v % 10

		if vup > 0 && long.Next == nil {
			long.Next = &ListNode{}
		}
		long = long.Next
	}

	return longroot
}

func main() {
	// result: [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// v1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	v1 := []int{2}
	v2 := []int{8, 9, 9}
	nodes := addTwoNumbers(createList(v1), createList(v2))
	fmt.Println(printList(nodes))
}
