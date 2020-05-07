package leetcode_go

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	q := [][]int{[]int{src, 0}}
	step := 0
	m := make(map[int][][]int)
	for _, flight := range flights {
		m[flight[0]] = append(m[flight[0]], []int{flight[1], flight[2]})
	}
	res := math.MaxInt32
	var curLen int
	for len(q) != 0 {
		curLen = len(q)
		for i := 0; i < curLen; i++ {
			if q[i][0] == dst {
				res = min(res, q[i][1])
			}
			if nexts, ok := m[q[i][0]]; ok {
				for _, next := range nexts {
					if q[i][1]+next[1] > res {
						continue
					}
					q = append(q, []int{next[0], q[i][1] + next[1]})
				}
			}
		}
		q = q[curLen:]
		if step > K {
			break
		}
		step++
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
