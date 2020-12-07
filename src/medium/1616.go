package main

func checkPalindromeFormation(a string, b string) bool {

	//四种情况分类讨论
	n := len(a)

	flag, ok := false, true
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		if !flag {
			if a[i] != b[j] {
				flag = !flag
				i--
			}
		} else {
			if a[i] != a[j] {
				ok = false
				break
			}
		}
	}

	if ok {
		return ok
	}

	flag, ok = false, true
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		if !flag {
			if b[i] != a[j] {
				flag = !flag
				i--
			}
		} else {
			if b[i] != b[j] {
				ok = false
				break
			}
		}
	}

	if ok {
		return ok
	}

	flag, ok = false, true
	for i := n - 1; i >= (n+1)/2; i-- {
		j := n - 1 - i
		if !flag {
			if a[i] != b[j] {
				flag = !flag
				i++
			}
		} else {
			if a[i] != a[j] {
				ok = false
				break
			}
		}
	}

	if ok {
		return ok
	}

	flag, ok = false, true
	for i := n - 1; i >= (n+1)/2; i-- {
		j := n - 1 - i
		if !flag {
			if b[i] != a[j] {
				flag = !flag
				i++
			}
		} else {
			if b[i] != b[j] {
				ok = false
				break
			}
		}
	}

	return ok
}
