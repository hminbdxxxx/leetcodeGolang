package leetcode_go

func deepestLeavesSum(root *TreeNode) int {
	res := 0
	q := []*TreeNode{root}
	var curLen int
	for len(q) > 0 {
		curLen = len(q)
		nextLev := []*TreeNode{}
		res = 0
		for i := 0; i < curLen; i++ {
			res += q[i].Val
			if q[i].Left != nil {
				nextLev = append(nextLev, q[i].Left)
			}
			if q[i].Right != nil {
				nextLev = append(nextLev, q[i].Right)
			}
		}
		if len(nextLev) == 0 {
			return res
		} else {
			q = append(q, nextLev...)
			q = q[curLen:]
		}
	}
	return res
}
