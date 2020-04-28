package leetcode_go

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	vals := []string{}
	fmt.Println(node.Val)
	for node != nil {
		// fmt.Println(node.Val)
		vals = append(vals, strconv.Itoa(node.Val))
		node = node.Next
	}
	fmt.Println(len(vals))
	fmt.Println(vals)
	return strings.Join(vals, "->")
}

// Node ...
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
