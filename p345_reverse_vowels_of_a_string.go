package leetcode_go

func reverseVowels(s string) string {
	res := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		for !isVowel(s[i]) {
			i++
		}
		for !isVowel(s[j]) {
			j--
		}
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(s)
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}
