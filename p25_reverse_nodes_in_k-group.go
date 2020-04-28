package leetcode_go

func reverseKGroup(head *ListNode, k int) *ListNode {
	l, r := head, head
	for i := 0; i < k; i++ {
		if r == nil {
			return head
		}
		r = r.Next
	}
	newHead := reverseHelperP25(l, r)
	l.Next = reverseKGroup(r, k)
	return newHead
}

func reverseHelperP25(left, right *ListNode) *ListNode {
	cur := left
	var pre, temp *ListNode
	for cur != right {
		temp = cur.Next
		cur.Next = pre

		pre = cur
		cur = temp
	}
	return pre
}
