package leetcode_go

func largestRectangleArea(heights []int) int {
	var res int
	var w int
	heights = append(heights, 0)
	st := make([]int, 0, len(heights))
	for idx := 0; idx < len(heights); {
		if len(st) == 0 || heights[idx] >= heights[st[len(st)-1]] {
			st = append(st, idx)
			idx++
		} else {
			t := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				w = idx
			} else {
				w = idx - st[len(st)-1] - 1
			}
			res = max(res, heights[t]*w)
		}
	}
	return res
}
