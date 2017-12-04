package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralCoords(t *testing.T) {
	assert.Equal(t, coords{0, 0}, spiralCoords(1))
	assert.Equal(t, coords{2, 1}, spiralCoords(12))
	assert.Equal(t, coords{0, -2}, spiralCoords(23))
}

func TestSpiralsize(t *testing.T) {
	assert.Equal(t, 1, spiralSize(0))
	assert.Equal(t, 8, spiralSize(1))
	assert.Equal(t, 16, spiralSize(2))
	assert.Equal(t, 24, spiralSize(3))
}
