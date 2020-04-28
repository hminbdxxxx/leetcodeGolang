package leetcode_go

func findBottomLeftValue(root *TreeNode) int {
	res := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []*TreeNode{}
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			level = append(level, q[i])
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[curLen:]
		if len(q) == 0 {
			res = level[0].Val
		}
	}
	return res
}
