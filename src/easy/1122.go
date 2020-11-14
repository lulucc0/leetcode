package main

import (
	"sort"
)

type slice struct {
	Arr1 []int
	Arr2 map[int]int
}

func (p *slice) Len() int {
	return len(p.Arr1)
}

func (p *slice) Less(i, j int) bool {

	idx1, ok1 := p.Arr2[p.Arr1[i]]
	idx2, ok2 := p.Arr2[p.Arr1[j]]

	if ok1 != ok2 {
		return ok1
	}

	if ok1 {
		return idx1 < idx2
	}

	return p.Arr1[i] < p.Arr1[j]
}

func (p *slice) Swap(i, j int) {
	p.Arr1[i], p.Arr1[j] = p.Arr1[j], p.Arr1[i]
}

func relativeSortArray(arr1 []int, arr2 []int) []int {

	m := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		m[v] = i
	}

	s := slice{arr1, m}
	sort.Sort(&s)

	return s.Arr1
}
