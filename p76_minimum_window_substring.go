package leetcode_go

import (
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	needs := make(map[byte]int)
	for _, v := range []byte(t) {
		needs[v]++
	}

	left, right := 0, 0
	st, ed := right, left
	minLen := math.MaxInt32
	missing := len(t)
	for right <= len(s) {
		if _, ok := needs[s[right]]; ok {
			needs[s[right]]--
			if needs[s[right]] >= 0 {
				missing--
			}
		}
		right++

		for missing == 0 && left < right {
			if right-left < minLen {
				minLen = right - left
				st = left
				ed = right
			}
			if _, ok := needs[s[left]]; ok {
				needs[s[left]]++
				if needs[s[left]] > 0 {
					missing++
				}
			}
			left++
		}
	}
	if minLen > len(s) {
		return ""
	}
	return s[st:ed]
}
