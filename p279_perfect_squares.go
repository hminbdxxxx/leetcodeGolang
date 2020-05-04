package leetcode_go

func numSquares(n int) int {
	q := []int{n}
	var curLen int
	vis := make(map[int]bool)
	res := 0
	for len(q) > 0 {
		curLen = len(q)
		for i := 0; i < curLen; i++ {
			if q[i] == 0 {
				return res
			}
			for j := 1; j*j <= q[i]; j++ {
				if _, ok := vis[q[i]-j*j]; !ok {
					q = append(q, q[i]-j*j)
					vis[q[i]-j*j] = true
				}
			}
		}
		q = q[curLen:]
		res++
	}
	return res
}
