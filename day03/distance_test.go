package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	assert.Equal(t, 0, distance(coords{}))
	assert.Equal(t, 3, distance(coords{2, 1}))
	assert.Equal(t, 2, distance(coords{0, -2}))
	assert.Equal(t, 31, distance(spiralCoords(1024)))
}
