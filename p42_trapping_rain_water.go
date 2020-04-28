package leetcode_go

func trap(height []int) int {
	l, r := 0, len(height)-1
	var minH, res int
	for l < r {
		minH = min(height[l], height[r])
		if minH == height[l] {
			l++
			for ; l < r && height[l] < minH; l++ {
				res += minH - height[l]
			}
		} else {
			r--
			for ; l < r && height[r] < minH; r-- {
				res += minH - height[r]
			}
		}
	}
	return res
}
