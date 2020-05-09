package leetcode_go

func characterReplacement(s string, k int) int {
	start := 0
	bs := []byte(s)
	cnt := make(map[byte]int)
	maxCnt := 0
	res := 0
	for i := 0; i < len(bs); i++ {
		cnt[bs[i]]++
		maxCnt = max(maxCnt, cnt[bs[i]])
		for i-start+1-maxCnt > k {
			cnt[bs[start]]--
			start++
		}
		res = max(res, i-start+1)
	}
	return res
}
