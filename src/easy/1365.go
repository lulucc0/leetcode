package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	sortNums := make([]int, len(nums))
	for i := range nums {
		sortNums[i] = nums[i]
	}

	sort.Ints(sortNums)
	m := make(map[int]int, len(nums))
	for i := len(sortNums) - 1; i >= 0; i-- {
		m[sortNums[i]] = i
	}

	ret := make([]int, len(nums))
	for i, num := range nums {
		ret[i] = m[num]
	}

	return ret
}
