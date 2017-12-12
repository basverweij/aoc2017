package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	assert.Equal(t, 3, cubeCoordsFromMoves(ne, ne, ne).distance())
	assert.Equal(t, 0, cubeCoordsFromMoves(ne, ne, sw, sw).distance())
	assert.Equal(t, 2, cubeCoordsFromMoves(ne, ne, s, s).distance())
	assert.Equal(t, 3, cubeCoordsFromMoves(se, sw, se, sw, sw).distance())
}
