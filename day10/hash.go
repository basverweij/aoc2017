package day10

import "encoding/hex"

var lengthsSuffix = []int{17, 31, 73, 47, 23}

const rounds = 64

func Hash(size int, input []byte) string {
	lengths := make([]int, len(input)+len(lengthsSuffix))
	for i, b := range input {
		lengths[i] = int(b)
	}
	for i, l := range lengthsSuffix {
		lengths[len(input)+i] = l
	}

	l := newList(size)
	h := make([]byte, size>>4)

	pos, skip := 0, 0
	for i := 0; i < rounds; i++ {
		l, pos, skip = knotWithParams(l, pos, skip, lengths)
	}

	for i := range h {
		h[i] = byte(l[i<<4])
		for j := 1; j < 16; j++ {
			h[i] ^= byte(l[i<<4+j])
		}
	}

	return hex.EncodeToString(h)
}
