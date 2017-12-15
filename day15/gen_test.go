package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInputGenA = 65
	testInputGenB = 8921
)

func TestNewGen(t *testing.T) {
	g := newGen(factorGenA, testInputGenA)
	assert.NotNil(t, g)
}

func TestGenNext(t *testing.T) {
	g := newGen(factorGenA, testInputGenA)
	assert.Equal(t, 1092455, g.next())
	assert.Equal(t, 1181022009, g.next())
	assert.Equal(t, 245556042, g.next())
	assert.Equal(t, 1744312007, g.next())
	assert.Equal(t, 1352636452, g.next())

	g = newGen(factorGenB, testInputGenB)
	assert.Equal(t, 430625591, g.next())
	assert.Equal(t, 1233683848, g.next())
	assert.Equal(t, 1431495498, g.next())
	assert.Equal(t, 137874439, g.next())
	assert.Equal(t, 285222916, g.next())
}
