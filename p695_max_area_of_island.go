package leetcode_go

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfsP695(grid, m, n, i, j))
		}
	}
	return res
}

func dfsP695(grid [][]int, m, n, i, j int) int {
	dirs := [][]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	area := 1
	for _, dir := range dirs {
		area += dfsP695(grid, m, n, i+dir[0], j+dir[1])
	}
	return area
}
