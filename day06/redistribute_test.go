package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedistribute(t *testing.T) {
	assert.Equal(t, 5, redistribute([]int{0, 2, 7, 0}))
}

func TestMax(t *testing.T) {
	max, bank := max([]int{0, 2, 7, 0})

	assert.Equal(t, 2, bank)
	assert.Equal(t, 7, max)
}
