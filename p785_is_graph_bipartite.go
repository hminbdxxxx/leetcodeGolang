// https://leetcode.com/problems/is-graph-bipartite/

package leetcode_go

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if colors[i] != 0 {
			continue
		}
		colors[i] = 1
		q := []int{i}
		for len(q) > 0 {
			t := q[0]
			q = q[1:]
			for _, node := range graph[t] {
				if colors[node] == colors[t] {
					return false
				}
				if colors[node] == 0 {
					colors[node] = -1 * colors[t]
					q = append(q, node)
				}
			}
		}
	}
	return true
}
