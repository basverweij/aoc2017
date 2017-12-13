package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeverity(t *testing.T) {
	s, _ := severity(testInput, 0, false)
	assert.Equal(t, 24, s)
}

func TestClearPass(t *testing.T) {
	assert.Equal(t, 10, clearPass(testInput))
}
