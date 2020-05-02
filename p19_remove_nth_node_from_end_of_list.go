package leetcode_go

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slowPre := dummy
	fast, slow := head, head
	for i := 1; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slowPre = slowPre.Next
		slow = slow.Next
	}

	slowPre.Next = slow.Next
	slow.Next = nil
	return dummy.Next
}
