package leetcode_go

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx, maxNum := 0, nums[0]
	for idx, num := range nums {
		if num > maxNum {
			maxIdx = idx
			maxNum = num
		}
	}
	return &TreeNode{
		Val:   maxNum,
		Left:  constructMaximumBinaryTree(nums[:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
}
