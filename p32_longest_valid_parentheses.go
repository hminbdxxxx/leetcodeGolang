package leetcode_go

func longestValidParentheses(s string) int {
	stk := []int{}
	res := 0
	startIdx := 0
	for idx, char := range s {
		if char == '(' {
			stk = append(stk, idx)
		} else {
			if len(stk) == 0 {
				startIdx = idx + 1
			} else {
				stk = stk[:len(stk)-1]
				if len(stk) == 0 {
					res = max(res, idx-startIdx+1)
				} else {
					res = max(res, idx-stk[len(stk)-1])
				}
			}
		}
	}
	return res
}
