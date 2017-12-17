package main

type buf struct {
	d   []int
	len int
	pos int
}

func newBuf(capacity int) *buf {
	return &buf{make([]int, capacity), 1, 0}
}

func (b *buf) move(amount int) {
	b.pos = (b.pos + amount) % b.len
}

func (b *buf) insert(val int) {
	b.pos++

	copy(b.d[b.pos+1:], b.d[b.pos:b.len])
	b.d[b.pos] = val

	b.len++
}

func (b *buf) buf() []int {
	return b.d[:b.len]
}
