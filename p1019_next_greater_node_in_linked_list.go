package leetcode_go

type elemP1019 struct {
	num int
	idx int
}

func nextLargerNodes(head *ListNode) []int {
	l := linkListLen(head)
	stk := []elemP1019{}
	res := make([]int, l)
	i := 0
	for head != nil {
		if len(stk) != 0 && head.Val > stk[len(stk)-1].num {
			j := len(stk) - 1
			for ; j >= 0 && stk[j].num < head.Val; j-- {
				res[stk[j].idx] = head.Val
			}
			stk = stk[:j+1]
		}
		stk = append(stk, elemP1019{num: head.Val, idx: i})
		head = head.Next
		i++
	}
	return res
}

func linkListLen(head *ListNode) int {
	l := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		l++
	}
	return l
}
