package main

func getFolderNames(names []string) []string {

	nameHash := map[string]int{}
	for i := 0; i < len(names); i++ {
		name := names[i]
		if count := nameHash[name]; count == 0 {
			nameHash[name] = 1
		} else {
			for nameHash[name+"("+string(count)+")"] > 0 {
				count++
			}
			newName := name + "(" + string(count) + ")"
			nameHash[newName] = 1
			nameHash[name] = count + 1
			names[i] = newName
		}
	}

	return names
}
