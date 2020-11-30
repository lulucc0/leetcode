package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reorganizeString("vvvlo"))
}

type item767 struct {
	Char  rune
	Count int
}

func reorganizeString(S string) string {

	if len(S) <= 0 {
		return ""
	}

	m := map[rune]int{}
	for _, c := range S {
		m[c] = m[c] + 1
	}

	cl := make([]item767, 0, len(m))
	for c, count := range m {
		cl = append(cl, item767{c, count})
	}

	sort.Slice(cl, func(i, j int) bool {
		return cl[i].Count > cl[j].Count
	})

	mid := (len(S) + 1) / 2

	//最多的超过一半则不可以
	if cl[0].Count > mid {
		return ""
	}

	sl, sr := "", ""
	nums := 0

	for _, c := range cl {
		str := string(c.Char)
		for i := 0; i < c.Count; i++ {
			if nums < mid {
				sl += str
			} else {
				sr += str
			}
			nums++
		}
	}

	ret, i, j, flag := "", 0, 0, true

	for i < mid || j < (len(S)-mid) {
		if flag {
			ret += string(sl[i])
			i++
		} else {
			ret += string(sr[j])
			j++
		}
		flag = !flag
	}

	return ret
}
