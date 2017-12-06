package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedistribute(t *testing.T) {
	steps, loop := redistribute([]int{0, 2, 7, 0})

	assert.Equal(t, 5, steps)
	assert.Equal(t, 4, loop)
}

func TestMax(t *testing.T) {
	max, bank := max([]int{0, 2, 7, 0})

	assert.Equal(t, 2, bank)
	assert.Equal(t, 7, max)
}
