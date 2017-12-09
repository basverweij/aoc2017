package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	assert.Equal(t, 1, score(scan("{}")))
	assert.Equal(t, 6, score(scan("{{{}}}")))
	assert.Equal(t, 5, score(scan("{{},{}}")))
	assert.Equal(t, 16, score(scan("{{{},{},{{}}}}")))
	assert.Equal(t, 1, score(scan("{<{},{},{{}}>}")))
	assert.Equal(t, 1, score(scan("{<a>,<a>,<a>,<a>}")))
	assert.Equal(t, 9, score(scan("{{<a>},{<a>},{<a>},{<a>}}")))
	assert.Equal(t, 3, score(scan("{{<!>},{<!>},{<!>},{<a>}}")))
}
