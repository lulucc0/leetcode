package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	ret := 0
	shot := 0
	for i := 0; i < len(points); i++ {
		if i == 0 {
			ret++
			shot = points[i][1]
		} else {
			if points[i][0] <= shot {
				continue
			} else {
				ret++
				shot = points[i][1]
			}
		}
	}
	return ret
}
