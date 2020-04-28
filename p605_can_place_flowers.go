package leetcode_go

func canPlaceFlowers(flowerbed []int, n int) bool {
	var pre, next, cnt int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i == 0 {
			pre = 0
		} else {
			pre = flowerbed[i-1]
		}
		if i == len(flowerbed)-1 {
			next = 0
		} else {
			next = flowerbed[i+1]
		}
		if pre == 0 && next == 0 {
			flowerbed[i] = 1
			cnt++
		}
	}
	return cnt >= n
}
