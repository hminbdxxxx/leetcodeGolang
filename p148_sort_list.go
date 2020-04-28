package leetcode_go

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	return mergeP148(sortList(head), sortList(slow.Next))
}

func mergeP148(h1, h2 *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			tail.Next = h1
			tail = h1
			h1 = h1.Next
		} else {
			tail.Next = h2
			tail = h2
			h2 = h2.Next
		}
	}
	if h1 != nil {
		tail.Next = h1
	} else {
		tail.Next = h2
	}
	return dummy.Next
}
