package leetcode_go

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := toStackP445(l1)
	stack2 := toStackP445(l2)
	carry := 0
	var preNode *ListNode
	head := ListNode{}
	for i, j := len(stack1)-1, len(stack2)-1; i > 0 || j > 0 || carry > 0; {
		var val int
		if i > 0 && j > 0 {
			val = stack1[i] + stack2[j] + carry
			i--
			j--
		} else if i < 0 {
			val = stack2[j] + carry
			j--
		} else if j < 0 {
			val = stack1[i] + carry
			i--
		} else {
			val = carry
		}
		val, carry = val%10, val/10
		node := ListNode{Val: val, Next: preNode}
		preNode = &node
		head.Next = &node
	}
	return head.Next
}

func toStackP445(node *ListNode) []int {
	stack := []int{}
	for node != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}
	return stack
}
