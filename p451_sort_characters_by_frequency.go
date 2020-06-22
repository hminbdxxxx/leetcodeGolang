// https://leetcode.com/problems/sort-characters-by-frequency/

package leetcode_go

import (
	"sort"
)

type CharCnt struct {
	b byte
	c int
}

type ByCnt []CharCnt

func (a ByCnt) Len() int           { return len(a) }
func (a ByCnt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCnt) Less(i, j int) bool { return a[i].c > a[j].c }


func frequencySort(s string) string {
	bs := []byte(s)
	cnt := make(map[byte]int)
	for _, b := range bs {
		cnt[b]++
	}
	byCnt := ByCnt{}
	for k, v := range cnt {
		byCnt = append(byCnt, CharCnt{b: k, c: v})
	}
	sort.Sort(byCnt)
	bs = []byte{}
	for _, c := range byCnt {
		for i := 0; i < c.c; i++ {
			bs = append(bs, c.b)
		}
	}
	return string(bs)
}