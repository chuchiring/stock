package main

import "fmt"

/*
topic 25
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	//calc length
	nlen, node := 0, head
	for node != nil {
		nlen++
		node = node.Next
	}

	if k > nlen {
		return head
	}

	//create a array
	arr := make([]*ListNode, k)
	node = head
	// i := 0
	j := k - 1
	var root, priorTail *ListNode
	for node != nil {
		arr[j] = node

		//filled
		if j == 0 {
			j = k - 1

			//1->2->3->4->5
			if root == nil {
				root = arr[0]
			}

			tail := arr[0].Next
			for h := 1; h < k; h++ {
				arr[h-1].Next = arr[h]
			}
			arr[k-1].Next = tail

			node = tail

			if priorTail != nil {
				priorTail.Next = arr[0]
			}
			priorTail = arr[k-1]
		} else {
			j--
			node = node.Next
		}

	}

	return root
}

func main() {
	root := IntArrayToSingeLinkList([]int{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(SingleLinkListToIntArray(reverseKGroup(root, 2)))
}
