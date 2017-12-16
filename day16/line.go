package main

type line struct {
	// holds the positions of the dancers
	d []int
}

func newLine(size int) *line {
	d := make([]int, size)
	for i := range d {
		d[i] = i
	}

	return &line{d}
}

func (l *line) String() string {
	b := make([]byte, len(l.d))
	for i, d := range l.d {
		b[d] = 'a' + byte(i)
	}

	return string(b)
}
