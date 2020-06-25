// https://leetcode.com/problems/single-number-iii/

package leetcode_go

func singleNumberIII(nums []int) []int {
	xor := 0
	res := []int{0, 0}
	for _, n := range nums {
		xor = xor ^ n
	}
	xor = xor & -xor
	for _, n := range nums {
		if (n & xor) == 0 {
			res[0] = res[0] ^ n
		} else {
			res[1] = res[1] ^ n
		}
	}
	return res
}
