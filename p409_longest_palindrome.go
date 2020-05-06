package leetcode_go

func longestPalindrome(s string) int {
	cnt := make(map[rune]int)
	for _, r := range s {
		cnt[r]++
	}

	res := 0
	for _, c := range cnt {
		res += (c / 2) * 2
	}

	if res < len(s) {
		res++
	}
	return res
}
