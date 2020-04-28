package leetcode_go

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	level := 1
	q := []*TreeNode{root}
	for len(q) != 0 {
		if level == d-1 {
			for _, node := range q {
				tmpLeft := node.Left
				tmpRight := node.Right
				node.Left = &TreeNode{Val: v, Left: tmpLeft}
				node.Right = &TreeNode{Val: v, Right: tmpRight}
			}
			break
		} else {
			curLen := len(q)
			for i := 0; i < curLen; i++ {
				if q[i].Left != nil {
					q = append(q, q[i].Left)
				}
				if q[i].Right != nil {
					q = append(q, q[i].Right)
				}
			}
			q = q[curLen:]
			level++
		}
	}
	return root
}
