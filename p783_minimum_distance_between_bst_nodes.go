// https://leetcode.com/problems/minimum-distance-between-bst-nodes/

package leetcode_go

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt32
	pre := -1
	inOrderP783(root, &res, &pre)
	return res
}

func inOrderP783(root *TreeNode, mn *int, pre *int) {
	if root.Left != nil {
		inOrderP783(root.Left, mn, pre)
	}
	if *pre != -1 {
		*mn = min(*mn, root.Val-*pre)
	}
	*pre = root.Val
	if root.Right != nil {
		inOrderP783(root.Right, mn, pre)
	}
}
