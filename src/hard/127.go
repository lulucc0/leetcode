package main

//广度优先搜索
func ladderLength(beginWord string, endWord string, wordList []string) int {

	length := 2
	bfs := []int{}
	in := make([]bool, len(wordList))
	cgm := map[int][]int{}
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if canChange(wordList[i], wordList[j]) {
				cgm[i] = append(cgm[i], j)
				cgm[j] = append(cgm[j], i)
			}
		}

		if canChange(beginWord, wordList[i]) {
			bfs = append(bfs, i)
			in[i] = true
			if wordList[i] == endWord {
				return length
			}
		}
	}

	for len(bfs) > 0 {
		length++
		newBfs := []int{}
		for _, i := range bfs {
			for _, j := range cgm[i] {
				if !in[j] {
					newBfs = append(newBfs, j)
					in[j] = true
					if wordList[j] == endWord {
						return length
					}
				}
			}
		}
		bfs = newBfs
	}

	return 0
}

func canChange(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	num := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			num++
		}
	}

	return num == 1
}
