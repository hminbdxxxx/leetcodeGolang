package leetcode_go

func findTilt(root *TreeNode) int {
	res := 0
	helperP563(root, &res)
	return res
}

func helperP563(root *TreeNode, tiltSum *int) int {
	if root == nil {
		return 0
	}
	leftSum := helperP563(root.Left, tiltSum)
	rightSum := helperP563(root.Right, tiltSum)
	tilt := leftSum - rightSum
	if tilt < 0 {
		tilt = -1 * tilt
	}
	*tiltSum += tilt
	return leftSum + rightSum + root.Val
}
