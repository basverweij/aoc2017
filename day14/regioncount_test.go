package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionCount(t *testing.T) {
	g := newGrid(testInput)

	assert.Equal(t, 1242, regionCount(g))
}
