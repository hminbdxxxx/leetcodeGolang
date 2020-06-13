// https://leetcode.com/problems/intersection-of-two-arrays-ii/

package leetcode_go

func intersect(nums1 []int, nums2 []int) []int {
	cnt := make(map[int]int)
	res := []int{}
	for _, i := range nums1 {
		cnt[i]++
	}
	for _, i := range nums2 {
		if cnt[i] > 0 {
			cnt[i]--
			res = append(res, i)
		}
	}
	return res
}
