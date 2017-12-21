package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPartsCollisions = []*particle{
	{p: d3{x: -6}, v: d3{x: 3}},
	{p: d3{x: -4}, v: d3{x: 2}},
	{p: d3{x: -2}, v: d3{x: 1}},
	{p: d3{x: 3}, v: d3{x: -1}},
}

func TestBruteForce(t *testing.T) {
	_, particlesLeft := bruteForce(testPartsCollisions, true)

	assert.Equal(t, 1, particlesLeft)
}
