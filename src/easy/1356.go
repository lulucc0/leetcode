package main

import (
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {

	n1, n2 := 0, 0
	a, b := p[i], p[j]
	for a > 0 {
		n1 += a % 2
		a /= 2
	}
	for b > 0 {
		n2 += b % 2
		b /= 2
	}

	if n1 == n2 {
		return p[i] < p[j]
	}

	return n1 < n2
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortByBits(arr []int) []int {

	sort.Sort(IntSlice(arr))
	return arr
}
