package leetcode_go

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := [][]int{}
	var start, end int
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i][0] > B[j][1] {
			j++
		} else if A[i][1] < B[j][0] {
			i++
		} else {
			start = max(A[i][0], B[j][0])
			end = min(A[i][1], B[j][1])
			res = append(res, []int{start, end})
			if A[i][1] < B[j][1] {
				i++
			} else {
				j++
			}
		}
	}
	return res
}
