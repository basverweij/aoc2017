package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrongestBridge(t *testing.T) {
	s := testInput.strongestBridge()

	assert.Equal(t, 31, s)
}
