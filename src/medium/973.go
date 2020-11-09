package main

func kClosest(points [][]int, K int) [][]int {

	close := make([]int, len(points))
	//构建大顶堆
	for i := 0; i < K; i++ {
		close[i] = points[i][0]*points[i][0] + points[i][1]*points[i][1]
		pi := i
		for pi > 0 {
			j := (pi+1)/2 - 1
			if close[pi] > close[j] {
				close[pi], close[j] = close[j], close[pi]
				points[pi], points[j] = points[j], points[pi]
				pi = j
			} else {
				break
			}
		}
	}

	//小于大顶堆的跟，则替换跟进堆调整
	for i := K; i < len(points); i++ {
		close[i] = points[i][0]*points[i][0] + points[i][1]*points[i][1]
		if close[i] < close[0] {
			close[i], close[0] = close[0], close[i]
			points[i], points[0] = points[0], points[i]

			pi := 0
			for pi*2+1 < K {
				j := -1
				if close[pi] < close[pi*2+1] {
					j = pi*2 + 1
				}
				if pi*2+2 < K && close[pi*2+2] > close[pi] && close[pi*2+2] > close[pi*2+1] {
					j = pi*2 + 2
				}
				if j >= 0 {
					close[pi], close[j] = close[j], close[pi]
					points[pi], points[j] = points[j], points[pi]
					pi = j
				} else {
					break
				}
			}
		}
	}

	return points[0:K]
}
