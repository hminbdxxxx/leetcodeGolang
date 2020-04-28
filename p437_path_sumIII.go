package leetcode_go

func pathSum3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helperP437(root, sum) + pathSum3(root.Left, sum) + pathSum3(root.Right, sum)
}

func helperP437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	if root.Val == sum {
		cnt = 1
	}
	return cnt + helperP437(root.Left, sum-root.Val) + helperP437(root.Right, sum-root.Val)
}
