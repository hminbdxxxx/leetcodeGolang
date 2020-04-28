package leetcode_go

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	curPath := []int{}
	dfsP113(root, sum, curPath, &res)
	return res
}

func dfsP113(root *TreeNode, sum int, curPath []int, res *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		curPath = append(append([]int{}, curPath...), root.Val)
		*res = append(*res, curPath)
		return
	}
	curPath = append(append([]int{}, curPath...), root.Val)
	dfsP113(root.Left, sum-root.Val, curPath, res)
	dfsP113(root.Right, sum-root.Val, curPath, res)
}
