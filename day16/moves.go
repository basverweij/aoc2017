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
	cutoff := l.size - s.size

	d := l.dancers >> uint(cutoff*4)
	d |= l.dancers << uint(s.size*4) & l.mask

	l.dancers = d
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
	nameA, nameB := l.get(e.posA), l.get(e.posB)

	l.set(e.posA, nameB)
	l.set(e.posB, nameA)
}

type partner struct {
	nameA, nameB uint64
}

func newPartner(nameA, nameB byte) *partner {
	return &partner{uint64(nameA - 'a'), uint64(nameB - 'a')}
}

func (p *partner) String() string {
	return fmt.Sprintf("p%s/%s", string(p.nameA+'a'), string(p.nameB+'a'))
}

func (p *partner) perform(l *line) {
	posA, posB := -1, -1
	for i := 0; i < l.size; i++ {
		n := l.get(i)
		if n == p.nameA {
			posA = i
		} else if n == p.nameB {
			posB = i
		}

		if posA >= 0 && posB >= 0 {
			break
		}
	}

	l.set(posA, p.nameB)
	l.set(posB, p.nameA)
}
