package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `flqrgnkx`

func TestNewGrid(t *testing.T) {
	g := newGrid(testInput)
	assert.NotNil(t, g)
}

func TestGridRowHash(t *testing.T) {
	g := newGrid(testInput)

	cases := map[int]byte{
		0: 212,
		1: 85,
		2: 10,
		3: 173,
		4: 104,
		5: 201,
		6: 68,
		7: 214,
	}

	for row, expectedHash := range cases {
		actualHash := g.rowHash(row)

		assert.Equal(t, expectedHash, actualHash[0], fmt.Sprintf("row %d", row))
	}
}

func TestGridUsage(t *testing.T) {
	g := newGrid(testInput)

	assert.Equal(t, 8108, g.usage())
}
