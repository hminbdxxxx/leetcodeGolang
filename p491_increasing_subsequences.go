package leetcode_go

func FfindSubsequences(nums []int) [][]int {
	res := [][]int{}
	helperP491(nums, []int{}, 0, &res)
	return res
}

func helperP491(nums []int, curSeq []int, pos int, res *[][]int) {
	if len(curSeq) > 1 {
		*res = append(*res, curSeq)
	}
	numSet := make(map[int]bool)
	for i := pos; i < len(nums); i++ {
		if _, ok := numSet[nums[i]]; ok {
			continue
		}
		if len(curSeq) == 0 || nums[i] >= curSeq[len(curSeq)-1] {
			curSeq = append(append([]int{}, curSeq...), nums[i])
			numSet[nums[i]] = true
			helperP491(nums, curSeq, i+1, res)
			curSeq = curSeq[:len(curSeq)-1]
		}
	}
}
