package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	m := input(testInput)
	d := runDiagnostics(m)

	assert.Equal(t, 3, d)
}
