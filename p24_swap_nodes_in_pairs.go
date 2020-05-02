package leetcode_go

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		last := pre.Next.Next
		pre.Next.Next = last.Next
		last.Next = pre.Next
		pre.Next = last
		pre = last.Next
	}
	return dummy.Next
}
