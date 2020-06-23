// https://leetcode.com/problems/contiguous-array/

package leetcode_go

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	res := 0
	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}
		if pre, ok := m[sum]; ok {
			res = max(res, i-pre)
		} else {
			m[sum] = i
		}
	}
	return res
}
