package leetcode_go

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := []*Node{root}
	for len(q) > 0 {
		level := []*Node{}
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
		pre := level[0]
		for _, cur := range level[1:] {
			pre.Next = cur
			pre = cur
		}
	}
	return root
}
