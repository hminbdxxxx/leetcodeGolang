package leetcode_go

func plusOne(digits []int) []int {
	carry := 0
	last := len(digits) - 1
	digits[last] = digits[last] + 1
	digits[last], carry = digits[last]/10, digits[last]%10
	for j := last - 1; j >= 0 && carry > 0; j-- {
		digits[j] = digits[j] + carry
		digits[j], carry = digits[j]/10, digits[j]%10
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
