package leetcode_go

import (
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	digs := strings.Split(digits, "")
	helperP17([]byte{}, 0, digs, &res)
	return res
}

func helperP17(cur []byte, pos int, digits []string, res *[]string) {
	if pos == len(digits) {
		(*res) = append((*res), string(cur))
	}

	keys := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "qprs", "tuv", "wxyz",
	}
	num, _ := strconv.ParseInt(digits[pos], 10, 32)
	letters := []byte(keys[num])
	for _, letter := range letters {
		cur = append(append([]byte{}, cur...), letter)
		helperP17(cur, pos+1, digits, res)
		cur = cur[:len(cur)-1]
	}
}
