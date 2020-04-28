package leetcode_go

import "math"

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt32
	pre := -1
	helperP530(root, &pre, &res)
	return res
}

func helperP530(node *TreeNode, pre *int, res *int) {
	if node.Left != nil {
		helperP530(node.Left, pre, res)
	}
	if *pre != -1 {
		*res = min(*res, node.Val-(*pre))
	}
	*pre = node.Val
	if node.Right != nil {
		helperP530(node.Right, pre, res)
	}
}
