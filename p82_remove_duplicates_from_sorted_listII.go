package leetcode_go

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := ListNode{Val: -1, Next: head}
	pre := &dummy
	for pre.Next != nil {
		cur := pre.Next
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next != cur {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
