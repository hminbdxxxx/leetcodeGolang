// https://leetcode.com/problems/possible-bipartition/

package leetcode_go

func possibleBipartition(N int, dislikes [][]int) bool {
	graph := [][]int{}
	for i := 0; i <= N; i++ {
		graph = append(graph, []int{})
	}
	for _, line := range dislikes {
		graph[line[0]] = append(graph[line[0]], line[1])
		graph[line[1]] = append(graph[line[1]], line[0])
	}
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
