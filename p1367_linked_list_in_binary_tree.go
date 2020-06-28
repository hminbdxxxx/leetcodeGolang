// https://leetcode.com/problems/linked-list-in-binary-tree/

package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return helperP1367(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func helperP1367(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return head.Val == root.Val && (helperP1367(head.Next, root.Left) || helperP1367(head.Next, root.Right))
}
