package main

type line struct {
	d []byte
}

func newLine(size int) *line {
	d := make([]byte, size)
	for i := range d {
		d[i] = 'a' + byte(i)
	}

	return &line{d}
}
