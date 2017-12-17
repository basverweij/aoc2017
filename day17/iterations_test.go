package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterations(t *testing.T) {
	b := newBuf(2018)

	iterations(b, 2017, 3)
	b.move(1)

	assert.Equal(t, 638, b.d[b.pos])
}
