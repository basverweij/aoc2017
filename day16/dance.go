package main

func dance(l *line, moves []move) {
	for _, m := range moves {
		m.perform(l)
	}
}
