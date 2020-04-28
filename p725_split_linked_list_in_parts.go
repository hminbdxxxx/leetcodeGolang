package leetcode_go

func splitListToParts(root *ListNode, k int) []*ListNode {
	if root == nil {
		return []*ListNode{}
	}
	length := 1
	for node := root; node.Next != nil; node = node.Next {
		length++
	}
	size, mod := length/k, length%k

	res := make([]*ListNode, k)
	node := root
	for i := 0; node != nil && i < k; i++ {
		res[i] = node
		curSize := size
		if mod > 0 {
			curSize++
			mod--
		}
		for j := 0; j < curSize-1; j++ {
			node = node.Next
		}
		next := node.Next
		node.Next = nil
		node = next
	}
	return res
}
