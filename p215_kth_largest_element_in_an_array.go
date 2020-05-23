// https://leetcode.com/problems/kth-largest-element-in-an-array/

package leetcode_go

func findKthLargest(nums []int, k int) int {
	var res int
	heaptify(nums)
	for ; k > 0; k-- {
		res = pop(&nums)
	}
	return res
}

func pop(arr *[]int) int {
	top := (*arr)[0]
	(*arr)[0] = (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	down(*arr, 0)
	return top
}

func heaptify(arr []int) {
	for i := len(arr)/2-1; i >=0; i-- {
		down(arr, i)
	}
}

func down(arr []int, i0 int) {
	i := i0
	for {
		j1 := i * 2 + 1
		if j1 >= len(arr) {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < len(arr) && arr[j2] > arr[j1] {
			j = j2
		}
		if arr[i] > arr[j] {
			break
		}
		arr[j], arr[i] = arr[i], arr[j]
		i = j
	}
}