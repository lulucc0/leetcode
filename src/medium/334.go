package main

func partitionLabels(S string) []int {

	last := map[rune]int{}
	for i, c := range S {
		last[c] = i
	}

	ret := make([]int, 0, len(last))
	lastI := 0
	nums := 0
	for i, c := range S {
		if last[c] > lastI {
			lastI = last[c]
		}
		nums++
		if i == lastI {
			ret = append(ret, nums)
			lastI, nums = 0, 0
		}
	}

	return ret
}
