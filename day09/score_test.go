package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func s(g *group, gc int) int {
	return score(g)
}

func TestScore(t *testing.T) {
	assert.Equal(t, 1, s(scan("{}")))
	assert.Equal(t, 6, s(scan("{{{}}}")))
	assert.Equal(t, 5, s(scan("{{},{}}")))
	assert.Equal(t, 16, s(scan("{{{},{},{{}}}}")))
	assert.Equal(t, 1, s(scan("{<{},{},{{}}>}")))
	assert.Equal(t, 1, s(scan("{<a>,<a>,<a>,<a>}")))
	assert.Equal(t, 9, s(scan("{{<a>},{<a>},{<a>},{<a>}}")))
	assert.Equal(t, 3, s(scan("{{<!>},{<!>},{<!>},{<a>}}")))
}
