package leetcode_go

func isIsomorphic(s string, t string) bool {
	preIdxS, preIdxT := make(map[byte]int), make(map[byte]int)
	for idx := 0; idx < len(s); idx++ {
		idx1, ok1 := preIdxS[s[idx]]
		idx2, ok2 := preIdxT[t[idx]]
		if ok1 && ok2 && idx1 != idx2 {
			return false
		}
		if ok1 != ok2 {
			return false
		}
		preIdxS[s[idx]], preIdxT[t[idx]] = idx, idx
	}
	return true
}
