// https://leetcode.com/problems/valid-parenthesis-string/

package leetcode_go

func checkValidString(s string) bool {
	mn, mx := 0, 0
	for _, ch := range []byte(s) {
		if ch == '(' {
			mx++
			mn++
		} else if ch == ')' {
			mx--
			mn = max(mn-1, 0)
		} else {
			mx++
			mn = max(mn-1, 0)
		}
		if mx < 0 {
			return false
		}
	}
	return mn == 0
}
