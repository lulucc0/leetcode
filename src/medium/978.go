package main

func maxTurbulenceSize(arr []int) int {

	ret := make([]int, len(arr))
	max := 0
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			ret[i] = 1
		} else {
			if arr[i] == arr[i-1] {
				ret[i] = 1
			} else {
				if ret[i-1] == 1 {
					ret[i] = 2
				} else {
					if (arr[i]-arr[i-1])*(arr[i-1]-arr[i-2]) < 0 {
						ret[i] = ret[i-1] + 1
					} else {
						ret[i] = 2
					}
				}
			}
		}
		if ret[i] > max {
			max = ret[i]
		}
	}
	return max
}
