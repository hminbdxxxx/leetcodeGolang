package leetcode_go

func convertBST(root *TreeNode) *TreeNode {
	curSum := 0
	helperP538(root, &curSum)
	return root
}

func helperP538(node *TreeNode, curSum *int) {
	if node.Right != nil {
		helperP538(node.Right, curSum)
	}
	tmp := node.Val
	node.Val = node.Val + *curSum
	*curSum += tmp
	if node.Left != nil {
		helperP538(node.Left, curSum)
	}
}
