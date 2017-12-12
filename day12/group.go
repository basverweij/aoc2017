package main

func group(programs map[int][]int, program int) map[int]struct{} {
	seen := make(map[int]struct{}, len(programs))
	todo := make([]int, 1)
	todo[0] = program

	var p int
	for len(todo) > 0 {
		p = todo[0]
		todo = todo[1:]

		for _, id := range programs[p] {
			if _, found := seen[id]; found {
				// do not visit twice
				continue
			}

			todo = append(todo, id)
		}

		seen[p] = struct{}{}
	}

	return seen
}

func groupCount(programs map[int][]int) int {
	seen := make(map[int]struct{}, len(programs))

	groupCount := 0
	for id := range programs {
		if _, found := seen[id]; found {
			// part of an already counted group
			continue
		}

		g := group(programs, id)
		for id := range g {
			seen[id] = struct{}{}
		}

		groupCount++
	}

	return groupCount
}
