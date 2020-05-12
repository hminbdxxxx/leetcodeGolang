// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

package leetcode_go

func countSmaller(nums []int) []int {
	sorted := []int{}
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		idx := biInsert(&sorted, nums[i])
		res[i] = idx
	}
	return res
}

func biInsert(arr *[]int, num int) int {
	l, r := 0, len(*arr)
	for l < r {
		mid := (l + r) / 2
		if (*arr)[mid] >= num {
			r = mid
		} else {
			l = mid + 1
		}
	}
	*arr = append(*arr, 0)
	copy((*arr)[r+1:], (*arr)[r:])
	(*arr)[r] = num
	return r
}
