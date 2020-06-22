// https://leetcode.com/problems/k-closest-points-to-origin/

package leetcode_go

import (
	"container/heap"
	"math"
)

type PointHeap [][]int

func (h PointHeap) Len() int {
	return len(h)
}

func (h PointHeap) Less(i, j int) bool {
	return math.Sqrt(float64(h[i][0]*h[i][0] + h[i][1]*h[i][1])) < math.Sqrt(float64(h[j][0]*h[j][0] + h[j][1]*h[j][1]))
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(point interface{}) {
	*h = append(*h, point.([]int))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, K int) [][]int {
	ph := PointHeap(points)
	heap.Init(&ph)
	res := [][]int{}
	for i := 0; i < K; i++ {
		res = append(res, heap.Pop(&ph).([]int))
	}
	return res
}