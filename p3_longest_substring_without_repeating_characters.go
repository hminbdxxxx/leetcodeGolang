package leetcode_go

func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 1, 1
	chars := []byte(s)
	m := make(map[byte]int)
	for left < len(s) {
		c := chars[right]
		right++
		m[c]++
		for m[c] > 1 {
			m[chars[left]]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}
