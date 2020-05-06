package leetcode_go

func findLHS(nums []int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	res := 0
	for n, c := range cnt {
		if c2, ok := cnt[n-1]; ok {
			res = max(res, c+c2)
		}
	}
	return res
}
