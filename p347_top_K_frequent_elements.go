package leetcode_go

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for i := 0; i < len(bucket); i++ {
		bucket[i] = []int{}
	}

	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	res := []int{}
Loop:
	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) > 0 {
			for _, num := range bucket[i] {
				res = append(res, num)
				k--
				if k == 0 {
					break Loop
				}
			}
		}
	}
	return res
}
