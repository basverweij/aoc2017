package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitCount(t *testing.T) {
	assert.Equal(t, 9, bitCount([]byte{0xa, 0x0c, 0x20, 0x17}))
}
