package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("aa", "aa"))
}

func minWindow(s string, t string) string {

	tm := map[uint8]int{}
	for i := 0; i < len(t); i++ {
		tm[t[i]] = tm[t[i]] + 1
	}

	in := map[uint8]int{}
	ok := map[uint8]bool{}
	min, mini, minj := 0, 0, 0
	for i, j := -1, 0; j < len(s); j++ {
		c := s[j]
		if tm[c] > 0 {
			if i < 0 {
				i = j
			}
			in[c] = in[c] + 1
			if in[c] >= tm[c] {
				ok[c] = true
			}
			if in[c] > 1 {
				for ; tm[s[i]] == 0 || in[s[i]] > tm[s[i]]; i++ {
					in[s[i]] = in[s[i]] - 1
				}
			}
			if len(ok) == len(tm) {
				if min == 0 || (j-i+1) < min {
					min = j - i + 1
					mini, minj = i, j
				}
			}
		}
	}

	if min == 0 {
		return ""
	}
	return s[mini : minj+1]
}
