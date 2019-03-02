package main

//ListNode single link node
type ListNode struct {
	Val  int
	Next *ListNode
}

//IntArrayToSingeLinkList change int array to a listnode list with int value
func IntArrayToSingeLinkList(arr []int) *ListNode {
	nlen := len(arr)
	if nlen == 0 {
		return nil
	}

	root := &ListNode{}
	node := root
	for i, v := range arr {
		node.Val = v
		if i != nlen-1 {
			node.Next = &ListNode{}
			node = node.Next
		}
	}
	return root
}

//SingleLinkListToIntArray 1
func SingleLinkListToIntArray(arr *ListNode) []int {
	var result []int

	node := arr
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}
