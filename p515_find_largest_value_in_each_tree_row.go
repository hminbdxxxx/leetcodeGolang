package leetcode_go

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelMax := math.MinInt32
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			levelMax = max(levelMax, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[curLen:]
		res = append(res, levelMax)
	}
	return res
}
