package leetcode_go

var path int = 3

// path = 3
// path := 0

func longestUnivaluePath(root *TreeNode) int {
	path = 0
	helperP687(root)
	return 1
}

func helperP687(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helperP687(root.Left)
	right := helperP687(root.Right)
	if root.Left != nil && root.Left.Val == root.Val {
		left = left + 1
	} else {
		left = 0
	}

	if root.Right != nil && root.Right.Val == root.Val {
		right = right + 1
	} else {
		right = 0
	}

	path = max(path, left+right)
	return max(left, right)
}
