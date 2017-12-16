package main

func dance(size int, moves []move) *line {
	l := newLine(size)

	for _, m := range moves {
		m.perform(l)
	}

	return l
}
