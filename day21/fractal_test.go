package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFractalize(t *testing.T) {
	start, rules := testInput()

	fractal := fractalize(start, rules)
	assert.Equal(t, 4, fractal.size)
	assert.Equal(t,
		"#..#/..../..../#..#",
		fractal.String())
	assert.Equal(t, 4, fractal.countOn())

	fractal = fractalize(fractal, rules)
	assert.Equal(t, 6, fractal.size)
	assert.Equal(t,
		"##.##./#..#../....../##.##./#..#../......",
		fractal.String())
	assert.Equal(t, 12, fractal.countOn())
}
