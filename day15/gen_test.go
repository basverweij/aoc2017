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
	g := newGen(factorGenA, testInputGenA, 1)
	assert.NotNil(t, g)
}

func TestGenNext(t *testing.T) {
	g := newGen(factorGenA, testInputGenA, 1)
	assert.Equal(t, 1092455, g.next())
	assert.Equal(t, 1181022009, g.next())
	assert.Equal(t, 245556042, g.next())
	assert.Equal(t, 1744312007, g.next())
	assert.Equal(t, 1352636452, g.next())

	g = newGen(factorGenB, testInputGenB, 1)
	assert.Equal(t, 430625591, g.next())
	assert.Equal(t, 1233683848, g.next())
	assert.Equal(t, 1431495498, g.next())
	assert.Equal(t, 137874439, g.next())
	assert.Equal(t, 285222916, g.next())
}

func TestGenNextWithMultiples(t *testing.T) {
	g := newGen(factorGenA, testInputGenA, multiplesGenA)
	assert.Equal(t, 1352636452, g.next())
	assert.Equal(t, 1992081072, g.next())
	assert.Equal(t, 530830436, g.next())
	assert.Equal(t, 1980017072, g.next())
	assert.Equal(t, 740335192, g.next())

	g = newGen(factorGenB, testInputGenB, multiplesGenB)
	assert.Equal(t, 1233683848, g.next())
	assert.Equal(t, 862516352, g.next())
	assert.Equal(t, 1159784568, g.next())
	assert.Equal(t, 1616057672, g.next())
	assert.Equal(t, 412269392, g.next())
}
