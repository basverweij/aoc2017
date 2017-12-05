package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
}

func TestPuzzle1(t *testing.T) {
	assert.Equal(t, 0, puzzle1(1))
	assert.Equal(t, 3, puzzle1(12))
	assert.Equal(t, 2, puzzle1(23))
	assert.Equal(t, 31, puzzle1(1024))
}

func TestPuzzle2(t *testing.T) {
	assert.Equal(t, 2, puzzle2(1))
	assert.Equal(t, 23, puzzle2(12))
	assert.Equal(t, 25, puzzle2(23))
	assert.Equal(t, 806, puzzle2(747))
}
