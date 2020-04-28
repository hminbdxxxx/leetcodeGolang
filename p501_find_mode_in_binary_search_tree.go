package leetcode_go

var curCnt, maxCnt int
var pre *TreeNode

func findMode(root *TreeNode) []int {
	curCnt, maxCnt = 0, 0
	res := []int{}
	inOrderP501(root, &res)
	return res
}

func inOrderP501(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrderP501(root.Left, res)
	if pre != nil && pre.Val == root.Val {
		curCnt++
	} else {
		curCnt = 1
	}

	if curCnt > maxCnt {
		maxCnt = curCnt
		*res = []int{root.Val}
	} else if curCnt == maxCnt {
		*res = append(*res, root.Val)
	}
	pre = root
	inOrderP501(root.Right, res)

}
