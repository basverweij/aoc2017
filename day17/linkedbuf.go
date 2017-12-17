package main

type item struct {
	next *item
	val  int
}

type linkedbuf struct {
	zero    *item
	current *item
	len     int
}

func newLinkedBuf() *linkedbuf {
	z := &item{}
	z.next = z

	return &linkedbuf{z, z, 1}
}

func (b *linkedbuf) move(amount int) {
	for i := 0; i < amount; i++ {
		b.current = b.current.next
	}
}

func (b *linkedbuf) insert(val int) {
	i := &item{val: val, next: b.current.next}
	b.current.next = i
	b.current = i
	b.len++
}

func (b *linkedbuf) buf() []int {
	buf := make([]int, b.len)

	item := b.current
	for i := range buf {
		buf[i] = item.val
		item = item.next
	}

	return buf
}
