package leetcode_go

func numComponents(head *ListNode, G []int) int {
	res := 0
	set := make(map[int]bool)
	for _, n := range G {
		set[n] = true
	}
	for head != nil {
		if inSetP817(head.Val, set) && (head.Next == nil || !inSetP817(head.Next.Val, set)) {
			res++
		}
		head = head.Next
	}
	return res
}

func inSetP817(n int, set map[int]bool) bool {
	_, ok := set[n]
	return ok
}
