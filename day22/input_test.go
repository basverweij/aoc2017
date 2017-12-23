package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawTestInput = `..#
#..
...`

func TestParse(t *testing.T) {
	ps := parse(rawTestInput)

	assert.Equal(t, []position{{1, -1}, {-1, 0}}, ps)
}
