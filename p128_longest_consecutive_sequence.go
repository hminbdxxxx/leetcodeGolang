package leetcode_go

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	var left, right, sum, res int
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			left = m[num-1]
			right = m[num+1]
			sum = left + right + 1
			res = max(res, sum)
			m[num] = sum
			m[num-left] = sum
			m[num+right] = sum
		}
	}
	return res
}
