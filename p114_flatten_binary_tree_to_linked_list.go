package leetcode_go

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tempRight := root.Right
	root.Right = root.Left
	for ; root.Right != nil; root = root.Right {
	}
	root.Right = tempRight
}
