package main

func uniquePaths(m int, n int) int {

	if m <= 0 || n <= 0 {
		return 0
	}

	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j] + ret[i][j-1]
			}
		}
	}

	return ret[m-1][n-1]
}
