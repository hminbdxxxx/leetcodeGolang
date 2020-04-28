package leetcode_go

func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters)-1
	var mid int
	for l <= h {
		mid = (l + h) / 2
		if letters[mid] >= target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l < len(letters) {
		return letters[l]
	}
	return letters[0]
}
