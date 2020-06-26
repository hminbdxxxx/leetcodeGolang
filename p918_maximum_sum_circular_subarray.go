// https://leetcode.com/problems/maximum-sum-circular-subarray/

package leetcode_go

import (
	"math"
)

func maxSubarraySumCircular(A []int) int {
	sum := 0
	mn, mx := math.MaxInt32, math.MinInt32
	curMin, curMax := 0, 0
	for _, n := range A {
		sum += n
		curMin = min(curMin+n, n)
		curMax = max(curMax+n, n)
		mn = min(mn, curMin)
		mx = max(mx, curMax)
	}
	if sum == mn {
		return mx
	}
	return max(sum-mn, mx)
}
