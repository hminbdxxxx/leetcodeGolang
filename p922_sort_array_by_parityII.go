package leetcode_go

func sortArrayByParityII(A []int) []int {
	oddIdx, evenIdx := 1, 0
	for oddIdx <= len(A)-1 && evenIdx <= len(A)-2 {
		for oddIdx <= len(A)-1 && A[oddIdx]%2 == 1 {
			oddIdx += 2
		}
		for evenIdx <= len(A)-2 && A[evenIdx]%2 == 0 {
			evenIdx += 2
		}
		if oddIdx <= len(A)-1 && evenIdx <= len(A)-2 {
			A[oddIdx], A[evenIdx] = A[evenIdx], A[oddIdx]
		}
	}
	return A
}
