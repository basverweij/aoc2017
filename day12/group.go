package main

func groupSize(programs map[int][]int, program int) int {
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

	return len(seen)
}
