package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	assert.Equal(t, 0, puzzle1(1))
	assert.Equal(t, 3, puzzle1(12))
	assert.Equal(t, 2, puzzle1(23))
	assert.Equal(t, 31, puzzle1(1024))
}
