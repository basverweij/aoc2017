package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func c(c *cubecoords, md int) *cubecoords {
	return c
}

func TestDistance(t *testing.T) {
	assert.Equal(t, 3, c(cubeCoordsFromMoves(ne, ne, ne)).distance())
	assert.Equal(t, 0, c(cubeCoordsFromMoves(ne, ne, sw, sw)).distance())
	assert.Equal(t, 2, c(cubeCoordsFromMoves(ne, ne, s, s)).distance())
	assert.Equal(t, 3, c(cubeCoordsFromMoves(se, sw, se, sw, sw)).distance())
}
