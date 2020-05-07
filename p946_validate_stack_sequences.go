package leetcode_go

func validateStackSequences(pushed []int, popped []int) bool {
	pIdx := 0
	stk := []int{}
	for _, num := range popped {
		if len(stk) == 0 || stk[len(stk)-1] != num {
			if pIdx >= len(pushed) {
				return false
			}
			for ; pIdx < len(pushed); pIdx++ {
				if pushed[pIdx] == num {
					pIdx++
					break
				} else {
					stk = append(stk, pushed[pIdx])
				}
			}
		} else {
			stk = stk[:len(stk)-1]
		}
	}
	return true
}
