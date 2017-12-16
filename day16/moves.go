package main

import "fmt"

type move interface {
	perform(l *line)
}

type spin struct {
	size int
}

func newSpin(size int) *spin {
	return &spin{size}
}

func (s *spin) String() string {
	return fmt.Sprintf("s%d", s.size)
}

func (s *spin) perform(l *line) {
	n := len(l.d)
	l.d = append(l.d[n-s.size:], l.d[0:n-s.size]...)
}

type exchange struct {
	posA, posB int
}

func newExchange(posA, posB int) *exchange {
	return &exchange{posA, posB}
}

func (e *exchange) String() string {
	return fmt.Sprintf("x%d/%d", e.posA, e.posB)
}

func (e *exchange) perform(l *line) {
	l.d[e.posA], l.d[e.posB] = l.d[e.posB], l.d[e.posA]
}

type partner struct {
	nameA, nameB byte
}

func newPartner(nameA, nameB byte) *partner {
	return &partner{nameA, nameB}
}

func (p *partner) String() string {
	return fmt.Sprintf("p%s/%s", string(p.nameA), string(p.nameB))
}

func (p *partner) perform(l *line) {
	posA, posB := 0, 0
	for i, n := range l.d {
		if n == p.nameA {
			posA = i
		}

		if n == p.nameB {
			posB = i
		}
	}

	l.d[posA], l.d[posB] = l.d[posB], l.d[posA]
}
