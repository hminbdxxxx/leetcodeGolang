package leetcode_go

func findFrequentTreeSum(root *TreeNode) []int {
	res := []int{}
	m := make(map[int]int)
	cnt := 0
	postorderP508(root, &cnt, &m, &res)
	return res
}

func postorderP508(node *TreeNode, cnt *int, cntMap *map[int]int, res *[]int) int {
	if node == nil {
		return 0
	}
	leftSum := postorderP508(node.Left, cnt, cntMap, res)
	rightSum := postorderP508(node.Right, cnt, cntMap, res)
	curSum := leftSum + rightSum + node.Val
	(*cntMap)[curSum]++
	if (*cntMap)[curSum] >= *cnt {
		if (*cntMap)[curSum] > *cnt {
			*res = []int{}
		}
		*res = append(*res, curSum)
		*cnt = (*cntMap)[curSum]
	}
	return curSum
}
