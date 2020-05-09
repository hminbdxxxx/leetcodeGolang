package leetcode_go

func rob3(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	res := helperP337(root, &m)
	return res
}

func helperP337(node *TreeNode, memo *map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	if val, ok := (*memo)[node]; ok {
		return val
	}
	val := 0
	if node.Left != nil {
		val += helperP337(node.Left.Left, memo) + helperP337(node.Left.Right, memo)
	}
	if node.Right != nil {
		val += helperP337(node.Right.Left, memo) + helperP337(node.Right.Right, memo)
	}
	res := max(node.Val+val, helperP337(node.Left, memo)+helperP337(node.Right, memo))
	(*memo)[node] = res
	return res
}
