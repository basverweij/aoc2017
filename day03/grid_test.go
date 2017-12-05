package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	g := newGrid(5, 2)

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			g.set(x, y, 3*(y+1)+x+1)
		}
	}

	t.Log(g)

	assert.Equal(t, 36, g.sumNeighbours(0, 0))

	assert.Equal(t, 8, g.sumNeighbours(-1, -1))
}
