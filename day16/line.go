package main

type line struct {
	size    int
	dancers uint64
	mask    uint64
}

func newLine(size int) *line {
	l := &line{size: size}

	l.mask = mask
	for i := 1; i < size; i++ {
		l.set(i, uint64(i))

		l.mask <<= 4
		l.mask |= mask
	}

	return l
}

func (l *line) String() string {
	b := make([]byte, l.size)

	for i := range b {
		b[i] = 'a' + byte(l.get(i))
	}

	return string(b)
}

const mask uint64 = 0xf

func (l *line) get(idx int) uint64 {
	return (l.dancers >> uint(idx*4)) & mask
}

func (l *line) set(idx int, val uint64) {
	l.dancers &= ^(mask << uint(idx*4))
	l.dancers |= (val << uint(idx*4))
}
