// https://leetcode.com/problems/range-sum-of-bst/

package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	inOrderP938(root, &res, L, R)
	return res
}

func inOrderP938(root *TreeNode, sum *int, L, R int) {
	if root.Left != nil {
		inOrderP938(root.Left, sum, L, R)
	}
	if root.Val >= L && root.Val <= R {
		(*sum) += root.Val
	}
	if root.Val <= R && root.Right != nil {
		inOrderP938(root.Right, sum, L, R)
	}
}
