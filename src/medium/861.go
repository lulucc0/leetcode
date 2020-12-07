package main

func matrixScore(A [][]int) int {

	//第一列必须是1   因此最优做法是  每一行都先翻转到第一位是1  然后再每列一次翻转得到最大值(当第一列全是0的情况下，先翻第一列以及把所有行都翻转 结果是一样的，因此不分类讨论)
	n := len(A)
	if n <= 0 {
		return n
	}

	m := len(A[0])
	bit := 1
	sum := 0

	//从最后一列即低位算起
	for j := m - 1; j >= 0; j-- {
		a, b := 0, 0
		for i := 0; i < n; i++ {
			if A[i][j]+A[i][0] == 1 { //01 或者 10
				b++
			} else { // 00 或者 11
				a++
			}
		}
		if a > b {
			sum += a * bit
		} else {
			sum += b * bit
		}
		bit *= 2
	}

	return sum
}
