package leetcode_go

func findLongestWord(s string, d []string) string {
	var res string
	for _, elem := range d {
		if len(elem) < len(res) || (len(elem) == len(res) && elem > res) {
			continue
		}
		if isSubP524(elem, s) {
			res = elem
		}
	}
	return res
}

func isSubP524(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return j == len(t)
}
