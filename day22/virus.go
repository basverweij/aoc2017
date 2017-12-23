package main

import (
	"bytes"
)

type position struct {
	x, y int
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

var moves = map[direction]position{
	up:    {0, -1},
	right: {1, 0},
	down:  {0, 1},
	left:  {-1, 0},
}

type level int

const (
	clean level = iota
	weakened
	infected
	flagged
)

type virus struct {
	pos        position
	dir        direction
	infected   map[position]level
	infections int
}

func newVirus(alreadyInfected []position) *virus {
	v := &virus{
		infected: make(map[position]level, len(alreadyInfected)),
	}

	for _, i := range alreadyInfected {
		v.infected[i] = infected
	}

	return v
}

func (v *virus) infectedCount() int {
	return len(v.infected)
}

func (v *virus) String(size int) string {
	mid := size / 2

	var buf bytes.Buffer
	for j := 0; j < size; j++ {
		if j > 0 {
			buf.WriteRune('\n')
		}

		if j-mid == v.pos.y && -mid == v.pos.x {
			buf.WriteRune('[')
		} else {
			buf.WriteRune(' ')
		}

		for i := 0; i < size; i++ {
			switch v.infected[position{i - mid, j - mid}] {
			case clean:
				buf.WriteRune('.')
			case weakened:
				buf.WriteRune('W')
			case infected:
				buf.WriteRune('#')
			case flagged:
				buf.WriteRune('F')
			}

			if j-mid == v.pos.y && i-mid == v.pos.x {
				buf.WriteRune(']')
			} else if j-mid == v.pos.y && i-mid+1 == v.pos.x {
				buf.WriteRune('[')
			} else {
				buf.WriteRune(' ')
			}
		}
	}

	return buf.String()
}

func (v *virus) burst() {
	l := v.infected[v.pos]

	if l == clean {
		// turn left
		v.dir = (v.dir + 3) % 4
	} else if l == infected {
		// turn right
		v.dir = (v.dir + 1) % 4
	} else if l == flagged {
		// reverse
		v.dir = (v.dir + 2) % 4
	}

	// modify
	l = (l + 1) % 4

	if l == clean {
		// clean
		delete(v.infected, v.pos)
	} else {
		v.infected[v.pos] = l

		if l == infected {
			v.infections++
		}
	}

	// move
	v.pos.x += moves[v.dir].x
	v.pos.y += moves[v.dir].y
}
