package leetcode_go

var maxPath int

func diameterOfBinaryTree(root *TreeNode) int {
	maxPath = 0
	singlePathP543(root)
	return maxPath
}

func singlePathP543(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	leftSinglePath := singlePathP543(root.Left)
	rightSinglePath := singlePathP543(root.Right)
	maxPath = max(maxPath, leftSinglePath+rightSinglePath+1)
	return max(leftSinglePath, rightSinglePath) + 1
}
