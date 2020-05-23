// https://leetcode.com/problems/binary-tree-postorder-traversal/

package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	path := []int{}
	stack := []*TreeNode{}
	cur := root
	var last *TreeNode = nil
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		if cur.Right == nil || last == cur.Right {
			path = append(path, cur.Val)
			stack = stack[:len(stack)-1]
			last = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return path
    
}