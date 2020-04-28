package leetcode_go

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	cached := make([][][]*TreeNode, n)
	for i := range cached {
		cached[i] = make([][]*TreeNode, n)
	}
	res := p95Helper(1, n, cached)
	return res
}

func p95Helper(st, ed int, cached [][][]*TreeNode) []*TreeNode {
	if st > ed {
		return []*TreeNode{nil}
	}
	if len(cached[st-1][ed-1]) > 0 {
		return cached[st-1][ed-1]
	}

	res := []*TreeNode{}
	for i := st; i <= ed; i++ {
		left, right := p95Helper(st, i-1, cached), p95Helper(i+1, ed, cached)
		for _, a := range left {
			for _, b := range right {
				node := TreeNode{
					Val:   i,
					Left:  a,
					Right: b,
				}
				res = append(res, &node)
			}
		}
	}
	cached[st-1][ed-1] = res
	return res
}
