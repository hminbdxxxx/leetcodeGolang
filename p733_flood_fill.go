// https://leetcode.com/problems/flood-fill/

package leetcode_go

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	vis := [][]bool{}
	for i := 0; i < len(image); i++ {
		tmp := []bool{}
		for j := 0; j < len(image[0]); j++ {
			tmp = append(tmp, false)
		}
		vis = append(vis, tmp)
	}
	helperP733(image, sr, sc, newColor, vis)
	return image
}

func helperP733(image [][]int, sr int, sc int, newColor int, vis [][]bool) {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	pre := image[sr][sc]
	image[sr][sc] = newColor
	for _, dir := range dirs {
		nr, nc := sr+dir[0], sc+dir[1]
		if nr >= 0 && nr < len(image) && nc >= 0 && nc < len(image[0]) && image[nr][nc] == pre && !vis[nr][nc] {
			vis[nr][nc] = true
			helperP733(image, nr, nc, newColor, vis)
		}
	}

}
