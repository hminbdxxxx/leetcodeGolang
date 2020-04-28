package leetcode_go

func insert(intervals [][]int, newInterval []int) [][]int {
	cur := 0
	res := [][]int{}
	for ; cur < len(intervals) && intervals[cur][1] < newInterval[0]; cur++ {
		res = append(res, intervals[cur])
	}

	for ; cur < len(intervals) && intervals[cur][0] < newInterval[1]; cur++ {
		newInterval[0] = min(intervals[cur][0], newInterval[0])
		newInterval[1] = max(intervals[cur][1], newInterval[1])
	}
	res = append(res, newInterval)

	for ; cur < len(intervals); cur++ {
		res = append(res, intervals[cur])
	}
	return res
}
