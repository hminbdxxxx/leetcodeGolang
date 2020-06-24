// https://leetcode.com/problems/find-the-town-judge/

package leetcode_go

func findJudge(N int, trust [][]int) int {
	indegree := make(map[int]int)
	outdegree := make(map[int]int)
	for _, line := range trust {
		indegree[line[1]]++
		outdegree[line[0]]++
	}
	for i := 1; i <= N; i++ {
		if indegree[i] == N-1 && outdegree[i] == 0 {
			return i
		}
	}
	return -1
}
