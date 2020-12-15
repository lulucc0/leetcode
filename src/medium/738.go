package main

import (
	"fmt"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(1234))
}

func monotoneIncreasingDigits(N int) int {

	nList := []int{}
	for N > 0 {
		nList = append(nList, N%10)
		N = N / 10
	}

	flag := false
	for i := len(nList) - 2; i >= 0; i-- {
		if flag {
			nList[i] = 9
		} else {
			if nList[i] < nList[i+1] {
				for j := i + 1; j < len(nList); j++ {
					if j == len(nList)-1 || nList[j] > nList[j+1] {
						flag = true
						nList[j] = nList[j] - 1
						i = j
						break
					}
				}
			}
		}
	}

	ret := 0
	bit := 1
	for i := 0; i < len(nList); i++ {
		ret += nList[i] * bit
		bit *= 10
	}

	return ret
}
