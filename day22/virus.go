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

type virus struct {
	pos        position
	dir        direction
	infected   map[position]struct{}
	infections int
}

func newVirus(alreadyInfected []position) *virus {
	v := &virus{
		infected: make(map[position]struct{}, len(alreadyInfected)),
	}

	for _, i := range alreadyInfected {
		v.infected[i] = struct{}{}
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
			if _, found := v.infected[position{i - mid, j - mid}]; found {
				buf.WriteRune('#')
			} else {
				buf.WriteRune('.')
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
	if _, found := v.infected[v.pos]; found {
		// turn right
		v.dir = (v.dir + 1) % 4

		// clean
		delete(v.infected, v.pos)
	} else {
		// turn left
		v.dir = (v.dir + 3) % 4

		// infect
		v.infected[v.pos] = struct{}{}
		v.infections++
	}

	// move
	v.pos.x += moves[v.dir].x
	v.pos.y += moves[v.dir].y
}
