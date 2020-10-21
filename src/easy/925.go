package main

func isLongPressedName(name string, typed string) bool {

	i, j, counti, countj := 0, 0, 0, 0

	for i < len(name) && j < len(typed) {
		for ; i < len(name)-1; i++ {
			if name[i] != name[i+1] {
				break
			} else {
				counti++
			}
		}

		for ; j < len(typed)-1; j++ {
			if typed[j] != typed[j+1] {
				break
			} else {
				countj++
			}
		}

		if name[i] != typed[j] || counti > countj {
			return false
		}

		i++
		j++
		counti, countj = 0, 0
	}

	return i == len(name) && j == len(typed)
}
