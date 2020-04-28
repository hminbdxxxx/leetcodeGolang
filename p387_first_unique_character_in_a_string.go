package leetcode_go

func firstUniqChar(s string) int {
	cnt := make(map[byte]int)
	sb := []byte(s)
	for _, c := range sb {
		cnt[c]++
	}

	for idx, c := range sb {
		if cnt[c] == 1 {
			return idx
		}
	}
	return -1
}
