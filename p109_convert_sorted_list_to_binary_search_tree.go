package leetcode_go

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	midPre, mid := findLinkedListMid(head)
	midPost := mid.Next
	midPre.Next = nil
	root := TreeNode{Val: mid.Val}
	if head != mid {
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(midPost)
	return &root
}

func findLinkedListMid(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	slowPre := slow
	for fast.Next != nil && fast.Next.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slowPre, slow
}
