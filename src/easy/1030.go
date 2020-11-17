package main

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	ret := make([][]int, 0, R*C)
	ret = append(ret, []int{r0, c0})

	in := make([][]bool, R)
	for i := range in {
		in[i] = make([]bool, C)
	}
	in[r0][c0] = true

	bfs := [][]int{{r0, c0}}
	for len(bfs) > 0 {
		nbfs := [][]int{}

		for _, point := range bfs {
			r, c := point[0], point[1]
			if r-1 >= 0 && !in[r-1][c] {
				in[r-1][c] = true
				nbfs = append(nbfs, []int{r - 1, c})
			}
			if c-1 >= 0 && !in[r][c-1] {
				in[r][c-1] = true
				nbfs = append(nbfs, []int{r, c - 1})
			}
			if r+1 < R && !in[r+1][c] {
				in[r+1][c] = true
				nbfs = append(nbfs, []int{r + 1, c})
			}
			if c+1 < C && !in[r][c+1] {
				in[r][c+1] = true
				nbfs = append(nbfs, []int{r, c + 1})
			}
		}

		if len(nbfs) > 0 {
			ret = append(ret, nbfs...)
		}

		bfs = nbfs
	}

	return ret
}
