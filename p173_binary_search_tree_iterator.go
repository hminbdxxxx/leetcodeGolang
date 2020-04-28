package leetcode_go

type BSTIterator struct {
	Stack []*TreeNode
}

func ConstructorP173(root *TreeNode) BSTIterator {
	iter := BSTIterator{Stack: []*TreeNode{}}
	for root != nil {
		iter.Stack = append(iter.Stack, root)
		root = root.Left
	}
	return iter
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	top := this.Stack[len(this.Stack)-1]
	res := top.Val
	this.Stack = this.Stack[:len(this.Stack)-1]
	if top.Right != nil {
		top = top.Right
		for top != nil {
			this.Stack = append(this.Stack, top)
			top = top.Left
		}
	}
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.Stack) == 0 {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
