// https://leetcode.com/problems/cousins-in-binary-tree/

package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	parentMap := make(map[int]int)
	levs := make(map[int]int)
	level := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
				parentMap[q[i].Left.Val] = q[i].Val
				levs[q[i].Left.Val] = level + 1
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
				parentMap[q[i].Right.Val] = q[i].Val
				levs[q[i].Right.Val] = level + 1
			}
		}
		level++
		q = q[curLen:]
	}
	return parentMap[x] != parentMap[y] && levs[x] == levs[y]
}
