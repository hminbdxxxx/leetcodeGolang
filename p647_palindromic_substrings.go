package leetcode_go

func OcountSubstrings(s string) int {
	isPalin := make([][]bool, len(s)+1)
	for i := range isPalin {
		arr := []bool{}
		for j := 0; j <= len(s); j++ {
			arr = append(arr, false)
		}
		isPalin[i] = arr
	}
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			isPalin[i][j] = (s[i] == s[j]) && (i-j <= 2 || isPalin[i-1][j+1])
			if isPalin[i][j] {
				res++
			}
		}
	}
	return res
}
