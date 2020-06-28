// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/

package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	helperP1022(root, &sum, []int{})
	return sum
}

func helperP1022(root *TreeNode, sum *int, path []int) {
	if root == nil {
		return
	}
	path = append(append([]int{}, path...), root.Val)
	if root.Left == nil && root.Right == nil {
		n := 0
		for i, j := len(path)-1, 1; i >= 0; i-- {
			n = n + path[i]*j
			j *= 2
		}
		*sum = *sum + n
	}
	if root.Left != nil {
		helperP1022(root.Left, sum, path)
	}
	if root.Right != nil {
		helperP1022(root.Right, sum, path)
	}
}
