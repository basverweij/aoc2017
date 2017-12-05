package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	assert.Equal(t, 2, fill(1))
	assert.Equal(t, 23, fill(12))
	assert.Equal(t, 25, fill(23))
	assert.Equal(t, 806, fill(747))
}
