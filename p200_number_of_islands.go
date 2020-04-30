package leetcode_go

func numIslands(grid [][]byte) int {
	vis := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		vis[i] = make([]bool, len(grid[0]))
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' || vis[i][j] {
				continue
			}
			helperP200(grid, i, j, &vis)
			res++
		}
	}
	return res
}

func helperP200(grid [][]byte, i, j int, visited *[][]bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' || (*visited)[i][j] {
		return
	}
	(*visited)[i][j] = true
	helperP200(grid, i+1, j, visited)
	helperP200(grid, i-1, j, visited)
	helperP200(grid, i, j+1, visited)
	helperP200(grid, i, j-1, visited)
}
