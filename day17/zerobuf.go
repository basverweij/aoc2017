package main

type zerobuf struct {
	zeroPos        int
	valueAfterZero int
	pos, len       int
}

func newZeroBuf() *zerobuf {
	return &zerobuf{len: 1}
}

func (b *zerobuf) move(amount int) {
	b.pos = (b.pos + amount) % b.len
}

func (b *zerobuf) insert(val int) {
	if b.pos == b.zeroPos {
		// we are only interested in the value inserted after zero
		b.valueAfterZero = val
	}

	b.pos++
	b.len++
}
