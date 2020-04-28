package leetcode_go

func preorderTraversal(root *TreeNode) []int {
	path := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			path = append(path, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = tmp.Right
		}
	}
	return path
}

func helperP144(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	helperP144(root.Left, path)
	helperP144(root.Right, path)
}
