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

	fCalc := func(node *ListNode) int {
		v := 0
		count := 1
		for node != nil {
			v = v + node.Val*count
			count = count * 10
			node = node.Next
		}
		return v
	}

	sum := fCalc(l1) + fCalc(l2)

	sumNode := &ListNode{}
	node := sumNode
	for {
		v := sum % 10
		node.Val = v

		sum = sum / 10
		if sum == 0 {
			break
		}

		node.Next = &ListNode{}
		node = node.Next
	}

	return sumNode
}

func main() {
	v1 := []int{2, 4, 3}
	v2 := []int{5, 6, 4}
	nodes := addTwoNumbers(createList(v1), createList(v2))
	fmt.Println(printList(nodes))
}
