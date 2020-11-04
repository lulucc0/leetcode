package main

func validMountainArray(A []int) bool {

	if len(A) < 3 || A[1] <= A[0] {
		return false
	}

	large := true
	for i := 2; i < len(A); i++ {
		if large {
			if A[i] == A[i-1] {
				return false
			} else if A[i] < A[i-1] {
				large = false
			}
		} else {
			if A[i] >= A[i-1] {
				return false
			}
		}
	}

	return !large
}
