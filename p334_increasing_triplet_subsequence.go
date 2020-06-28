// https://leetcode.com/problems/increasing-triplet-subsequence/

package leetcode_go

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	n1, n2 := math.MaxInt32, math.MaxInt32
	for _, n := range nums {
		if n <= n1 {
			n1 = n
		} else if n <= n2 {
			n2 = n
		} else {
			return true
		}
	}
	return false
}
