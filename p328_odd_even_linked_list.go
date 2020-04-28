package leetcode_go

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head

}
