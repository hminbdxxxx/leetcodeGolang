package leetcode_go

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := [][]int{}
	for len(q) > 0 {
		level := []int{}
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			cur := q[i]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		q = q[curLen:]
		res = append([][]int{level}, res...)
	}
	return res
}
