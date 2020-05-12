// https://leetcode.com/problems/max-consecutive-ones-iii/

package leetcode_go

func longestOnes(A []int, K int) int {
	res := 0
	l, r := 0, 0
	for ; r < len(A); r++ {
		if A[r] == 0 {
			K--
		}
		if K < 0 {
			if A[l] == 0 {
				K++
			}
			l++
		}
		res = r - l + 1
	}
	return res
}
