package leetcode_go

var pre99, q, p *TreeNode

func recoverTree(root *TreeNode) {
	pre99, q, p = nil, nil, nil
	traverse(root)
	q.Val, p.Val = p.Val, q.Val
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Left)
	if pre99 != nil && root.Val < pre.Val {
		if q == nil {
			q = pre99
		}
		p = root
	}
	pre99 = root
	traverse(root.Right)

}
