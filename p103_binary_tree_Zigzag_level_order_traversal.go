package leetcode_go

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := [][]int{}
	flag := true
	for len(q) > 0 {
		level := []int{}
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			cur := q[i]
			if flag {
				level = append(level, cur.Val)
			} else {
				level = append([]int{cur.Val}, level...)
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		flag = !flag
		q = q[curLen:]
		res = append(res, level)
	}
	return res
}
