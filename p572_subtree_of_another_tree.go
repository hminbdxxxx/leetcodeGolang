package leetcode_go

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isEqualP572(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isEqualP572(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		return isEqualP572(s.Left, t.Left) && isEqualP572(s.Right, t.Right)
	}
	return false
}
