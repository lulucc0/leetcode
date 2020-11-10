package main

func nextPermutation(nums []int) {

	done := false
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			done = true

			j := i
			for ; j < n-1; j++ {
				if nums[j] < nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				} else {
					break
				}
			}

			for j := j - 1; j > i; j-- {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}

			for k := i + 1; k < (i+1+n)/2; k++ {
				nums[k], nums[n-k+i] = nums[n-k+i], nums[k]
			}

			break
		}
	}

	if !done {
		for i := 0; i < n/2; i++ {
			j := n - i - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
