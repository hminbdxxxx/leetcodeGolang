// https://leetcode.com/problems/find-all-anagrams-in-a-string/

package leetcode_go

// import (
// 	"fmt"
// )

func findAnagrams(s string, p string) []int {
	res := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return res
	}
	sb, pb := []byte(s), []byte(p)
	m := make(map[byte]int)
	mp := make(map[byte]int)
	for i := 0; i < len(pb); i++ {
		m[sb[i]]++
		mp[pb[i]]++
	}
	if mapEqualP438(mp, m) {
		res = append(res, 0)
	}
	for i := len(pb); i < len(sb); i++ {
		m[sb[i]]++
		m[sb[i-len(pb)]]--
		if mapEqualP438(mp, m) {
			res = append(res, i-len(pb)+1)
		}
	}

	return res
}

func mapEqualP438(m1, m2 map[byte]int) bool {
	for k1, v1 := range m1 {
		if m2[k1] != v1 {
			return false
		}
	}
	return true
}
