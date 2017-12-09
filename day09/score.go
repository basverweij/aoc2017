package main

func score(root *group) int {
	return scoreFromLevel(root, 1)
}

func scoreFromLevel(group *group, level int) int {
	s := level

	for _, c := range group.children {
		s += scoreFromLevel(c, level+1)
	}

	return s
}
