package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {

	sslice := make([]string, 0, len(s))
	for _, c := range s {
		sslice = append(sslice, string(c))
	}

	sort.Strings(sslice)

	retList := []string{}
	nowRow := 0
	for i, c := range sslice {
		if i == 0 || c != sslice[i-1] {
			nowRow = 0
		} else {
			nowRow++
		}

		if len(retList) <= nowRow {
			retList = append(retList, c)
		} else {
			if nowRow%2 == 0 {
				retList[nowRow] = retList[nowRow] + c
			} else {
				retList[nowRow] = c + retList[nowRow]
			}
		}
	}

	return strings.Join(retList, "")
}
