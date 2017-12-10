package main

func knot(size int, lengths []int) list {
	l, _, _ := knotWithParams(newList(size), 0, 0, lengths)
	return l
}

func knotWithParams(l list, pos, skip int, lengths []int) (list, int, int) {
	size := len(l)

	for _, length := range lengths {
		l.Reverse(pos, length)

		pos = (pos + length + skip) % size
		skip++
	}

	return l, pos, skip
}
