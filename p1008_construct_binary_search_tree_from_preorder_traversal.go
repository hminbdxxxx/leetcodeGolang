// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := TreeNode{Val: preorder[0]}
	mid := 1
	for ; mid < len(preorder); mid++ {
		if preorder[mid] > preorder[0] {
			break
		}
	}
	root.Left = bstFromPreorder(preorder[1:mid])
	root.Right = bstFromPreorder(preorder[mid:])
	return &root
}
