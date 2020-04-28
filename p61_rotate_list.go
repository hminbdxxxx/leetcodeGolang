package leetcode_go

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	len := 1
	for cur := head; cur.Next != nil; cur = cur.Next {
		len++
	}
	slow, fast := head, head
	for i := 0; i < k%len && fast.Next != nil; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if slow.Next != nil {
		newHead := slow.Next
		slow.Next = nil
		fast.Next = head
		return newHead
	}
	return head
}
