package leetcode_go

func minDepth(root *TreeNode) int {
	level := 0
	if root == nil {
		return level
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		level++
		for i := 0; i < curLen; i++ {
			if q[i].Left == nil && q[i].Right == nil {
				return level
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[curLen:]
	}
	return level
}
