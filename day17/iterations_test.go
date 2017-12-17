package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterations(t *testing.T) {
	b := newBuf(2018)

	assert.Equal(t, 638, iterations(b, 2017, 3))
}
