package leetcode_go

func validPalindrome(s string) bool {
	t := []byte(s)
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		if t[i] != t[j] {
			return isPalindromeP680(t, i+1, j) || isPalindromeP680(t, i, j-1)
		}
	}
	return true
}

func isPalindromeP680(s []byte, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
