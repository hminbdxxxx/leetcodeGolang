package leetcode_go

import "math"

func countNodes(root *TreeNode) int {
	hleft, hright := 0, 0
	left, right := root, root
	for left != nil {
		left = left.Left
		hleft++
	}
	for right != nil {
		right = right.Right
		hright++
	}
	if hleft == hright {
		return int(math.Pow(float64(2), float64(hleft))) - 1
	}
	leftCnt := countNodes(root.Left)
	rightCnt := countNodes(root.Right)
	return leftCnt + rightCnt + 1
}
