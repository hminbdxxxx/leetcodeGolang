package leetcode_go

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	curPath := []string{}
	helperP257(root, curPath, &res)
	return res
}

func helperP257(node *TreeNode, curPath []string, res *[]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		curPath = append(append([]string{}, curPath...), strconv.Itoa(node.Val))
		path := strings.Join(curPath, "->")
		*res = append(*res, path)
		return
	}
	curPath = append(append([]string{}, curPath...), string(node.Val))
	helperP257(node.Left, curPath, res)
	helperP257(node.Right, curPath, res)
}
