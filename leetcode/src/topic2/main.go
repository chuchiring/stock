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

	root := &ListNode{}
	node := root

	for {
		var vup int
		if l1 != nil && l2 != nil {
			v := l1.Val + l2.Val + node.Val
			vup = v / 10
			v = v % 10

			node.Val = v
		} else if l1 != nil {
			v := node.Val + l1.Val
			vup = v / 10
			v = v % 10
			node.Val = v
		} else if l2 != nil {
			v := node.Val + l2.Val
			vup = v / 10
			v = v % 10
			node.Val = v
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && vup == 0 {
			break
		}

		node.Next = &ListNode{Val: vup}
		node = node.Next
	}

	return root
}

func main() {
	// result: [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// v1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	v1 := []int{1}
	v2 := []int{9, 9}
	nodes := addTwoNumbers(createList(v1), createList(v2))
	fmt.Println(printList(nodes))
}
