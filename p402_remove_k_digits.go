// https://leetcode.com/problems/remove-k-digits/

package leetcode_go

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	stk := []byte{}
	size := len(num) - k
	for _, n := range []byte(num) {
		for k > 0 && len(stk) > 0 && stk[len(stk)-1] > n {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, n)
	}
	if len(stk) > size {
		stk = stk[:size]
	}
	res := string(stk)
	res = strings.TrimLeft(res, "0")
	if len(res) == 0 {
		return "0"
	}
	return res
}
