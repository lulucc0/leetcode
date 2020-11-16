package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}

type slice [][]int

func (p slice) Len() int {
	return len(p)
}

func (p slice) Less(i, j int) bool {

	if p[i][0] == p[j][0] {
		return p[i][1] > p[j][1]
	}

	return p[i][0] < p[j][0]
}

func (p slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func reconstructQueue(people [][]int) [][]int {

	sort.Sort(slice(people))
	for i := len(people) - 2; i >= 0; i-- {
		nums := people[i][1]
		for j := i; j < i+nums; j++ {
			people[j], people[j+1] = people[j+1], people[j]
		}
	}
	return people
}
