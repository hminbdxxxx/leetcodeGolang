package leetcode_go

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	_, res = helper(root, res)
	return res
}

func helper(root *TreeNode, res int) (int, int) {
	if root == nil {
		return 0, math.MinInt64
	}
	leftSinglePath, leftRes := helper(root.Left, res)
	leftSinglePath = max(leftSinglePath, 0)
	rightSinglePath, rightRes := helper(root.Right, res)
	rightSinglePath = max(rightSinglePath, 0)
	res = max(rightRes, max(leftRes, leftSinglePath+rightSinglePath+root.Val))
	return max(leftSinglePath, rightSinglePath) + root.Val, res
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
