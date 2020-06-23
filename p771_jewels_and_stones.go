// https://leetcode.com/problems/jewels-and-stones/

package leetcode_go

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)
	for _, j := range J {
		m[j] = true
	}
	res := 0
	for _, s := range S {
		if _, ok := m[s]; ok {
			res++
		}
	}
	return res
}
