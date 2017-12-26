package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentStart(t *testing.T) {
	c := component{1, 2}
	assert.Equal(t, 1, c.start(fromA))
	assert.Equal(t, 2, c.start(fromB))
}

func TestComponentEnd(t *testing.T) {
	c := component{1, 2}
	assert.Equal(t, 2, c.end(fromA))
	assert.Equal(t, 1, c.end(fromB))
}

func TestComponentStrength(t *testing.T) {
	c := component{1, 2}
	assert.Equal(t, 3, c.strength())
}
