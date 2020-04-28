package leetcode_go

func twoSumP167(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	var sum int
	for l < r {
		sum = numbers[l] + numbers[r]
		if sum == target {
			return []int{l, r}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
