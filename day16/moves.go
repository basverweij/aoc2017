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
	cutoff := len(l.d) - s.size

	for i, idx := range l.d {
		if idx < cutoff {
			l.d[i] += s.size
		} else {
			l.d[i] -= cutoff
		}
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
	idxA, idxB := 0, 0
	for i, pos := range l.d {
		if pos == e.posA {
			idxA = i
		}

		if pos == e.posB {
			idxB = i
		}
	}

	l.d[idxA], l.d[idxB] = l.d[idxB], l.d[idxA]
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
	l.d[p.nameA], l.d[p.nameB] = l.d[p.nameB], l.d[p.nameA]
}
