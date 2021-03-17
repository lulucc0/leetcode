package main

func maximumSwap(num int) int {

	if num < 0 {
		return num
	}

	numList := []int{}
	lmaxList := []int{}
	rmaxList := []int{}
	for num > 0 {
		numList = append(numList, num%10)
		num /= 10
		if len(numList) == 1 {
			rmaxList = append(rmaxList, 0)
			lmaxList = append(lmaxList, 0)
		} else {
			lmax := len(numList) - 1
			rmax := len(numList) - 1
			if numList[rmax] < numList[rmaxList[rmax-1]] {
				rmax = rmaxList[rmax-1]
			}
			rmaxList = append(rmaxList, rmax)

			if numList[lmax] <= numList[lmaxList[lmax-1]] {
				lmax = lmaxList[lmax-1]
			}
			lmaxList = append(lmaxList, lmax)
		}
	}

	for i := len(numList) - 1; i > 0; i-- {
		if rmaxList[i] != i {
			j := lmaxList[i]
			numList[i], numList[j] = numList[j], numList[i]
			break
		}
	}

	b := 1
	ret := 0
	for i := 0; i < len(numList); i++ {
		ret += b * numList[i]
		b *= 10
	}
	return ret
}
