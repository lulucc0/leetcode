package main

func moveZeroes(nums []int) {

	i, j, n := 0, 0, len(nums)
	for {
		for i < n && nums[i] != 0 {
			i++
		}
		for j < n && (nums[j] == 0 || j <= i) {
			j++
		}

		if i >= n || j >= n {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j++
	}
}
