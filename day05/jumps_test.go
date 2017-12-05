package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepsToExit(t *testing.T) {
	js := jumps{0, 3, 0, 1, -3}

	assert.Equal(t, 5, js.StepsToExit())
}

func TestStepsToExitWithDecrease(t *testing.T) {
	js := jumps{0, 3, 0, 1, -3}

	assert.Equal(t, 10, js.StepsToExitWithDecrease())
}
