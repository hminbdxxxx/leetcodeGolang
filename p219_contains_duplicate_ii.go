// https://leetcode.com/problems/contains-duplicate-ii/
package leetcode_go

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if len(m) > k {
			delete(m, nums[i-k-1])
		}
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = i
	}
	return false
}
