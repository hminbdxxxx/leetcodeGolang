package leetcode_go

import "strings"

func simplifyPath(path string) string {
	ps := strings.Split(path, "/")
	st := make([]string, 0, len(ps))
	for _, c := range ps {
		if c == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else if c != "." && c != "" {
			st = append(st, c)
		}
	}
	res := "/" + strings.Join(st, "/")
	return res
}
