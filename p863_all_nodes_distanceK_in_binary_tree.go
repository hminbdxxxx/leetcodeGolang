package leetcode_go


func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	father := make(map[*TreeNode]*TreeNode)
	helperP863(root, nil, father)
	res := []int{}
	q := []*TreeNode{target}
	visited := make(map[*TreeNode]bool)
	visited[target] = true
	step := 0
	for len(q) != 0 {
		if step == K {
			for _, node := range q {
				res = append(res, node.Val)
			}
			return res
		}
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			if q[i].Left != nil && !visited[q[i].Left] {
				q = append(q, q[i].Left)
				visited[q[i].Left] = true
			}
			if q[i].Right != nil && !visited[q[i].Right] {
				q = append(q, q[i].Right)
				visited[q[i].Right] = true
			}
			if node, ok := father[q[i]]; ok && !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
		step++
		q = q[curLen:]
	}
	return res
}

func helperP863(node, father *TreeNode, m map[*TreeNode]*TreeNode) {
	if father != nil {
		m[node] = father
	}
	if node.Left != nil {
		helperP863(node.Left, node, m)
	}
	if node.Right != nil {
		helperP863(node.Right, node, m)
	}
}
