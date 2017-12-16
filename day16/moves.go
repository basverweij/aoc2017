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
	cutoff := len(l.dancers) - s.size

	copy(l.dancersSwap[:s.size], l.dancers[cutoff:])
	copy(l.dancersSwap[s.size:], l.dancers[:cutoff])

	l.dancers, l.dancersSwap = l.dancersSwap, l.dancers

	for i, d := range l.dancers {
		d.pos = i
	}
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
	// switch positions
	l.dancers[e.posA].pos, l.dancers[e.posB].pos = l.dancers[e.posB].pos, l.dancers[e.posA].pos

	// switch dancers
	l.dancers[e.posA], l.dancers[e.posB] = l.dancers[e.posB], l.dancers[e.posA]
}

type partner struct {
	nameA, nameB int
}

func newPartner(nameA, nameB byte) *partner {
	return &partner{int(nameA - 'a'), int(nameB - 'a')}
}

func (p *partner) String() string {
	return fmt.Sprintf("p%s/%s", string(p.nameA+'a'), string(p.nameB+'a'))
}

func (p *partner) perform(l *line) {
	// get positions
	posA := l.byName[p.nameA].pos
	posB := l.byName[p.nameB].pos

	// switch positions
	l.dancers[posA].pos, l.dancers[posB].pos = l.dancers[posB].pos, l.dancers[posA].pos

	// switch dancers
	l.dancers[posA], l.dancers[posB] = l.dancers[posB], l.dancers[posA]
}
