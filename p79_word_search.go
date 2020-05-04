package leetcode_go

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	vis := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			vis[i] = append(vis[i], false)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helperP79(board, []byte(word), i, j, 0, &vis) {
				return true
			}
		}
	}
	return false
}

func helperP79(board [][]byte, word []byte, i, j int, curLen int, vis *[][]bool) bool {
	if curLen == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || (*vis)[i][j] || board[i][j] != word[curLen] {
		return false
	}

	(*vis)[i][j] = true
	res := helperP79(board, word, i+1, j, curLen+1, vis) || helperP79(board, word, i-1, j, curLen+1, vis) || helperP79(board, word, i, j+1, curLen+1, vis) || helperP79(board, word, i, j-1, curLen+1, vis)
	(*vis)[i][j] = false
	return res
}
