// https://leetcode.com/problems/subarray-product-less-than-k/

package leetcode_go

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	i, j := 0, 0
	prod := 1
	res := 0
	for ; j < len(nums); j++ {
		prod *= nums[j]
		for prod >= k && i < len(nums) {
			prod /= nums[i]
			i++
		}
		res += (j - i + 1)
	}
	return res
}
