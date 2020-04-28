package leetcode_go

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
			if i == curLen-1 {
				res = append(res, q[i].Val)
			}
		}
		q = q[curLen:]
	}
	return res
}
