package leetcode_go

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := TreeNode{Val: pre[0]}
	pre = pre[1:]
	post = post[:len(post)-1]

	if len(pre) > 0 {
		preBegin := pre[0]
		var postMid int
		for idx, val := range post {
			if val == preBegin {
				postMid = idx
			}
		}
		root.Left = constructFromPrePost(pre[:postMid+1], post[:postMid+1])
		root.Right = constructFromPrePost(pre[postMid+1:], post[postMid+1:])
	}
	return &root
}
