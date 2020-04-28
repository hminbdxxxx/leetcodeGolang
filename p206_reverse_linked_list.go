package leetcode_go

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre, temp *ListNode
	for cur != nil {
		temp = cur.Next
		cur.Next = pre

		pre = cur
		cur = temp
	}
	return pre
}
