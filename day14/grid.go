package main

import (
	"encoding/hex"
	"fmt"

	"github.com/basverweij/aoc2017/day10"
)

const rowHashFormat = "%s-%d"

type grid struct {
	key string
}

func newGrid(key string) *grid {
	return &grid{key}
}

func (g *grid) rowHash(row int) []byte {
	h, _ := hex.DecodeString(day10.Hash(256, []byte(fmt.Sprintf(rowHashFormat, g.key, row))))

	return h
}

func (g *grid) usage() int {
	n := 0

	for i := 0; i < 128; i++ {
		n += bitCount(g.rowHash(i))
	}

	return n
}
