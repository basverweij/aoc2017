package main

func knot(size int, lengths []int) list {
	l := newList(size)

	pos, skip := 0, 0
	for _, length := range lengths {
		l.Reverse(pos, length)

		pos = (pos + length + skip) % size
		skip++
	}

	return l
}
