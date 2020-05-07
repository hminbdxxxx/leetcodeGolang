package leetcode_go

type MyStack struct {
	q []int
}

/** Initialize your data structure here. */
func MyStackConstructor() MyStack {
	return MyStack{q: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	for i := 0; i < len(this.q)-1; i++ {
		poped := this.q[0]
		this.q = append(this.q[1:], poped)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	poped := this.q[0]
	this.q = this.q[1:]
	return poped
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
