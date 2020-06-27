// https://leetcode.com/problems/leaf-similar-trees/

package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1, l2 := []int{}, []int{}
	helperP872(root1, &l1)
	helperP872(root2, &l2)
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func helperP872(root *TreeNode, leafSeq *[]int) {
	if root.Left == nil && root.Right == nil {
		*leafSeq = append(*leafSeq, root.Val)
		return
	}
	if root.Left != nil {
		helperP872(root.Left, leafSeq)
	}
	if root.Right != nil {
		helperP872(root.Right, leafSeq)
	}
}
