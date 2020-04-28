package leetcode_go

func checkInclusion(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	needs := make(map[byte]int)
	for _, b := range b1 {
		needs[b]++
	}
	l, r := 0, 0
	cnt := make(map[byte]int)
	for ; r < len(b2); r++ {
		cnt[b2[r]]++
		for mapContainP567(needs, cnt) {
			if mapEqualP567(needs, cnt) {
				return true
			}
			cnt[b2[l]]--
			if cnt[b2[l]] == 0 {
				delete(cnt, b2[l])
			}
			l++
		}
	}
	return false
}

func mapContainP567(m1, m2 map[byte]int) bool {
	for k1, v1 := range m1 {
		if m2[k1] < v1 {
			return false
		}
	}
	return true
}

func mapEqualP567(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if m2[k1] != v1 {
			return false
		}
	}
	return true
}
