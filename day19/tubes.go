package main

const (
	verti byte = '|'
	horiz byte = '-'
	cross byte = '+'
)

func isLetter(t byte) bool {
	return t >= 'A' && t <= 'Z'
}

func follow(tubes [][]byte) (string, int) {
	s := ""
	n := 0

	pos := findStart(tubes)
	dir := down

	for dir != none {
		n++
		pos.move(dir)

		t := tubes[pos.y][pos.x]

		if isLetter(t) {
			// add found letter, direction does not change
			s += string(t)
			continue
		}

		if t == cross {
			// cross changes direction from up/down to left/right and vice versa
			if dir == down || dir == up {
				if tubes[pos.y][pos.x+1] == horiz {
					dir = right
				} else if tubes[pos.y][pos.x-1] == horiz {
					dir = left
				} else if isLetter(tubes[pos.y][pos.x+1]) {
					dir = right
				} else if isLetter(tubes[pos.y][pos.x-1]) {
					dir = left
				} else {
					dir = none
				}
			} else {
				if tubes[pos.y+1][pos.x] == verti {
					dir = down
				} else if tubes[pos.y-1][pos.x] == verti {
					dir = up
				} else if isLetter(tubes[pos.y+1][pos.x]) {
					dir = down
				} else if isLetter(tubes[pos.y-1][pos.x]) {
					dir = up
				} else {
					dir = none
				}
			}
			continue
		}

		if t == verti || t == horiz {
			// either same direction, or a crossing of another line
			continue
		}

		dir = none
	}

	return s, n
}

func findStart(tubes [][]byte) *position {
	for i, t := range tubes[0] {
		if t == verti {
			return &position{i, 0}
		}
	}

	return nil
}
