package leetcode_go

func moveUpLockI(s string, i int) string {
	bs := []byte(s)
	if bs[i] == '9' {
		bs[i] = '0'
	} else {
		bs[i]++
	}
	return string(bs)
}

func moveDownLockI(s string, i int) string {
	bs := []byte(s)
	if bs[i] == '0' {
		bs[i] = '9'
	} else {
		bs[i]--
	}
	return string(bs)
}

func openLock(deadends []string, target string) int {
	q := []string{"0000"}
	step := 0
	visited := make(map[string]bool)
	for len(q) > 0 {
		curLen := len(q)
	elemLoop:
		for i := 0; i < curLen; i++ {
			if q[i] == target {
				return step
			}
			for _, dead := range deadends {
				if q[i] == dead {
					continue elemLoop
				}
			}
			for pos := 0; pos < 4; pos++ {
				up := moveUpLockI(q[i], pos)
				if _, ok := visited[up]; !ok {
					q = append(q, up)
					visited[up] = true
				}
				down := moveDownLockI(q[i], pos)
				if _, ok := visited[down]; !ok {
					q = append(q, down)
					visited[down] = true
				}

			}
		}
		q = q[curLen:]
		step++
	}
	return -1
}
