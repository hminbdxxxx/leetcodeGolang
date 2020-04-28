package leetcode_go

func dailyTemperatures(T []int) []int {
	stk := []int{}
	res := make([]int, len(T))
	for idx, temp := range T {
		for len(stk) > 0 && temp > T[stk[len(stk)-1]] {
			t := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res[t] = idx - t
		}
		stk = append(stk, idx)
	}
	return res
}
