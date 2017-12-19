package main

type direction int

const (
	none direction = iota
	down
	up
	right
	left
)

type position struct {
	x, y int
}

func (p *position) move(dir direction) {
	switch dir {
	case down:
		p.y++
	case up:
		p.y--
	case right:
		p.x++
	case left:
		p.x--
	}
}
