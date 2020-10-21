package main

func totalNQueens(n int) int {

	stack := make([][2]int, 0, n)

	if n <= 0 {
		return 0
	}

	var next = func(p [2]int) ([2]int, bool) {
		np := p
		ok := true

		if np[1] < n-1 {
			np[1] = np[1] + 1
		} else {
			for {
				if len(stack) == 0 {
					ok = false
					break
				}
				np = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				if np[1] < n-1 {
					np[1] = np[1] + 1
					break
				}
			}
		}

		return np, ok
	}

	point := [2]int{0, 0}
	sum := 0
	for {
		ok := true
		for _, old := range stack {
			if old[0] == point[0] || old[1] == point[1] || point[0]-old[0] == point[1]-old[1] || point[0]-old[0]+point[1]-old[1] == 0 { //判断皇后是否互相冲突
				ok = false
				break
			}
		}

		var have bool
		if !ok {
			point, have = next(point)
			if !have {
				break
			}
		} else {
			if point[0] == n-1 {
				sum++
				point, have = next(point)
				if !have {
					break
				}
			} else {
				stack = append(stack, [2]int{point[0], point[1]})
				point[0], point[1] = point[0]+1, 0
			}
		}
	}

	return sum
}
