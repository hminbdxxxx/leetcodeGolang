package leetcode_go

var memoP894 map[int][]*TreeNode = make(map[int][]*TreeNode)

func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return []*TreeNode{}
	}
	if N == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	if r, ok := memoP894[N]; ok {
		return r
	}
	res := []*TreeNode{}
	for i := 1; i < N; i += 2 {
		left, right := allPossibleFBT(i), allPossibleFBT(N-i-1)
		for _, l := range left {
			for _, r := range right {
				tmp := &TreeNode{Val: 0, Left: l, Right: r}
				res = append(res, tmp)
			}
		}
	}
	memoP894[N] = res
	return res
}
