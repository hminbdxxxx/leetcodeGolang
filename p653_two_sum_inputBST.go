package leetcode_go

func findTarget(root *TreeNode, k int) bool {
	nums := []int{}
	inOrderP653(root, &nums)
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == k {
			return true
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inOrderP653(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrderP653(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrderP653(root.Right, nums)
}
