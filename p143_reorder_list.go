package leetcode_go

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	cur := head
	stack := []*ListNode{}
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	midIdx := len(stack) / 2
	cur = head
	var curNext *ListNode
	for idx := len(stack) - 1; idx > midIdx; idx-- {
		t := stack[idx]
		stack = stack[:len(stack)-1]
		curNext = cur.Next
		cur.Next = t
		t.Next = curNext
		cur = curNext
	}
	stack[len(stack)-1].Next = nil
}
