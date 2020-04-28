package leetcode_go

func widthOfBinaryTree(root *TreeNode) int {
	type NodePos struct {
		node *TreeNode
		pos  int
	}

	res := 1
	q := []*NodePos{&NodePos{root, 1}}
	var curLen int
	for len(q) > 0 {
		curLen = len(q)
		left := q[0].pos
		right := left
		for i := 0; i < curLen; i++ {
			node := q[i].node
			right = q[i].pos
			if node.Left != nil {
				q = append(q, &NodePos{node.Left, right * 2})
			}
			if node.Right != nil {
				q = append(q, &NodePos{node.Right, right*2 + 1})
			}
		}
		q = q[curLen:]
		res = max(res, right-left+1)
	}
	return res
}
