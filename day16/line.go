package main

type dancer struct {
	name int
	pos  int
}

type line struct {
	byName      []*dancer
	dancers     []*dancer
	dancersSwap []*dancer
}

func newLine(size int) *line {
	l := &line{
		dancers:     make([]*dancer, size),
		dancersSwap: make([]*dancer, size),
		byName:      make([]*dancer, size),
	}

	for i := size - 1; i >= 0; i-- {
		l.dancers[i] = &dancer{name: i, pos: i}
		l.byName[i] = l.dancers[i]
	}

	return l
}

func (l *line) String() string {
	b := make([]byte, len(l.dancers))

	for i, d := range l.dancers {
		b[i] = 'a' + byte(d.name)
	}

	return string(b)
}
