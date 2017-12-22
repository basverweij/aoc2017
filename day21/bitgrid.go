package main

import (
	"fmt"
	"math/big"
	"strings"
)

type bitgrid struct {
	size int
	mask []*big.Int
}

func newBitgrid(size int) *bitgrid {
	b := &bitgrid{
		size: size,
		mask: make([]*big.Int, size),
	}

	for i := range b.mask {
		b.mask[i] = new(big.Int)
	}

	return b
}

func (b *bitgrid) String() string {
	s := make([]string, len(b.mask))

	for i, m := range b.mask {
		b := fmt.Sprintf(
			fmt.Sprintf("%%0%db", b.size),
			reverse(m, b.size))

		s[i] = strings.Replace(strings.Replace(b, "0", ".", -1), "1", "#", -1)
	}

	return strings.Join(s, "/")
}

var bigint1 = big.NewInt(1)

func (b *bitgrid) set(x, y int, mask bool) {
	var m uint
	if mask {
		m = 1
	}

	b.mask[y].SetBit(b.mask[y], x, m)
}

func (b *bitgrid) get(x, y int) bool {
	return b.mask[y].Bit(x) == 1
}

func (b *bitgrid) key(x, y, size int) string {
	mask := new(big.Int)
	mask.Lsh(bigint1, uint(size))
	mask.Sub(mask, bigint1)

	key := new(big.Int)
	for i := y + size - 1; i >= y; i-- {
		if i < y+size-1 {
			key.Lsh(key, uint(size))
		}

		m := new(big.Int)
		m.Rsh(b.mask[i], uint(x))
		m.And(m, mask)

		key.Or(key, m)
	}

	return key.Text(16)
}

func (b *bitgrid) countOn() int {
	c := 0
	for j := 0; j < b.size; j++ {
		c += bitCount(b.mask[j])
	}

	return c
}

func (b *bitgrid) copyFrom(from *bitgrid, x, y int) {
	fromMask := new(big.Int)
	fromMask.Lsh(bigint1, uint(from.size))
	fromMask.Sub(fromMask, bigint1)
	fromMask.Lsh(fromMask, uint(x))

	mask := new(big.Int)
	mask.Lsh(bigint1, uint(b.size))
	mask.Sub(mask, bigint1)
	mask.Xor(mask, fromMask)

	for j := 0; j < from.size; j++ {
		m := new(big.Int)
		m.Lsh(from.mask[j], uint(x))

		b.mask[y+j].And(b.mask[y+j], mask)
		b.mask[y+j].Or(b.mask[y+j], m)
	}
}
