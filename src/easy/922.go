package main

func sortArrayByParityII(A []int) []int {

	i, j := 0, 1
	for {
		for i < len(A) && A[i]%2 == 0 {
			i += 2
		}

		for j < len(A) && A[j]%2 == 1 {
			j += 2
		}

		if i >= len(A) || j >= len(A) {
			break
		}

		A[i], A[j] = A[j], A[i]
	}
	return A
}
