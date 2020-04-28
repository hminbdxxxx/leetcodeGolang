package leetcode_go

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p, q := headA, headB
	for p != q {
		if p.Next == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q.Next == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}
