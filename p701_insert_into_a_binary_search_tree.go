package leetcode_go

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root.Val = val
		return root
	}
	if root.Val <= val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Left, val)
		}
	}
	return root
}
