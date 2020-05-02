package leetcode_go

type intPair struct {
	x, y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	dirs := []intPair{
		intPair{0, 1},
		intPair{0, -1},
		intPair{1, 0},
		intPair{1, 1},
		intPair{1, -1},
		intPair{-1, 0},
		intPair{-1, 1},
		intPair{-1, -1},
	}
	m, n := len(grid), len(grid[0])
	q := []intPair{intPair{0, 0}}
	pathLen := 0
	for len(q) != 0 {
		pathLen++
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			if grid[q[i].x][q[i].y] == 1 {
				continue
			}
			if q[i].x == m-1 && q[i].y == n-1 {
				return pathLen
			}
			grid[q[i].x][q[i].y] = 1
			for _, dir := range dirs {
				newPos := intPair{q[i].x + dir.x, q[i].y + dir.y}
				if newPos.x < 0 || newPos.x >= m || newPos.y < 0 || newPos.y >= n || grid[newPos.x][newPos.y] == 1 {
					continue
				}
				q = append(q, newPos)
			}
		}
		q = q[curLen:]
	}
	return -1
}
