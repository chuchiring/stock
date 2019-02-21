package main

//TreeNode is a tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameNode(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil && left.Val == right.Val {
		return isSameNode(left.Left, right.Right) && isSameNode(left.Right, right.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSameNode(root.Left, root.Right)
}
