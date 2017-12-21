package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

type bitgrid struct {
	size int
	mask []uint
}

func newBitgrid(size int) *bitgrid {
	return &bitgrid{
		size: size,
		mask: make([]uint, size),
	}
}

func (b *bitgrid) String() string {
	s := make([]string, len(b.mask))

	for i, m := range b.mask {
		b := fmt.Sprintf(
			fmt.Sprintf("%%0%ds", b.size),
			strconv.FormatUint(bits.Reverse64(uint64(m))>>uint(64-b.size), 2))

		s[i] = strings.Replace(strings.Replace(b, "0", ".", -1), "1", "#", -1)
	}

	return strings.Join(s, "/")
}

const uint1 = uint(1)

func (b *bitgrid) set(x, y int, mask bool) {
	if mask {
		b.mask[y] |= uint1 << uint(x)
	} else {
		b.mask[y] &= ^(uint1 << uint(x))
	}
}

func (b *bitgrid) get(x, y int) bool {
	return (b.mask[y]>>uint(x))&uint1 == uint1
}

func (b *bitgrid) key(x, y, size int) uint {
	var key uint
	mask := (uint1 << uint(size)) - uint1

	for i := y + size - 1; i >= y; i-- {
		if i < y+size-1 {
			key <<= uint(size)
		}

		key |= (b.mask[i] >> uint(x)) & mask
	}

	return key
}

func (b *bitgrid) countOn() int {
	c := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.get(i, j) {
				c++
			}
		}
	}

	return c
}

func (b *bitgrid) copyFrom(from *bitgrid, x, y int) {
	mask := ^(((uint1 << uint(from.size)) - uint1) << uint(x))

	for j := 0; j < from.size; j++ {
		b.mask[y+j] &= mask
		b.mask[y+j] |= from.mask[j] << uint(x)
	}
}
