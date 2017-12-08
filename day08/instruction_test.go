package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInstructionFromSource(t *testing.T) {
	i, err := newInstructionFromSource("a inc -1 if b > -2")
	assert.Nil(t, err)

	assert.Equal(t, &instruction{"a", true, -1, "b", ">", -2}, i)
}
