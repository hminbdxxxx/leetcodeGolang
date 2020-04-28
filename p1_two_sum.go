package leetcode_go

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if anotherIdx, ok := m[num]; ok {
			return []int{idx, anotherIdx}
		}
		m[target-num] = idx
	}
	return []int{}
}
