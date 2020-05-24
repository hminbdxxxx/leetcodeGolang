// https://leetcode.com/problems/word-break-ii/

package leetcode_go

import (
	"strings"
)

func WwordBreak(s string, wordDict []string) []string {
	memo := make(map[string][]string)
	return backtrackP140(s, wordDict, &memo)
}

func backtrackP140(s string, wordDict []string, memo *map[string][]string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	if r, ok := (*memo)[s]; ok {
		return r
	}

	res := []string{}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			trimmed := strings.TrimPrefix(s, word)
			rem := backtrackP140(trimmed, wordDict, memo)
			for _, r := range rem {
				if len(r) == 0 {
					res = append(res, word)
				} else {
					res = append(res, word+" "+r)
				}
			}
		}
	}
	(*memo)[s] = res
	return res
}
