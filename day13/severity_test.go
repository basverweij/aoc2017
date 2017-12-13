package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeverity(t *testing.T) {
	assert.Equal(t, 24, severity(testInput))
}
