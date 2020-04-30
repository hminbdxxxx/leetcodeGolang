package leetcode_go

func flattenP430(root *NodeWithChild) *NodeWithChild {
	cur := root
	for cur.Next != nil {
		if cur.Child != nil {
			curNext := cur.Next
			nextLev := cur.Child
			for nextLev.Next != nil {
				nextLev = nextLev.Next
			}
			cur.Next = cur.Child
			cur.Next.Prev = cur
			cur.Child = nil
			nextLev.Next = curNext
			if curNext != nil {
				curNext.Prev = nextLev
			}
		}
		cur = cur.Next
	}
	return root
}
