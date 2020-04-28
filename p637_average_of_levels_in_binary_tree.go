package leetcode_go

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	res := []float64{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		level := []int{}
		for i := 0; i < curLen; i++ {
			level = append(level, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[curLen:]
		avg := avg(level)
		res = append(res, avg)
	}
	return res
}

func avg(nums []int) float64 {
	sum := 0
	if len(nums) == 0 {
		return 0
	}

	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}
