package main

func removeDuplicates(S string) string {

	stack := make([]uint8, 0, len(S))

	for i := 0; i < len(S); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != S[i] {
			stack = append(stack, S[i])
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}

	return string(stack)
}
