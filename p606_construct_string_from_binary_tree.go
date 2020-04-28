package leetcode_go

import "strconv"

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	left := tree2str(t.Left)
	right := tree2str(t.Right)

	if left == "" && right == "" {
		return strconv.Itoa(t.Val)
	}
	if right == "" {
		return strconv.Itoa(t.Val) + "(" + left + ")"
	}
	return strconv.Itoa(t.Val) + "(" + left + ")" + "(" + right + ")"
}
