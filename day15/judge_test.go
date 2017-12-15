package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudge(t *testing.T) {
	genA := newGen(factorGenA, testInputGenA, 1)
	genB := newGen(factorGenB, testInputGenB, 1)
	assert.Equal(t, 1, judge(genA, genB, 5))

	genA = newGen(factorGenA, testInputGenA, 1)
	genB = newGen(factorGenB, testInputGenB, 1)
	assert.Equal(t, 588, judge(genA, genB, 40000000))

	genA = newGen(factorGenA, testInputGenA, multiplesGenA)
	genB = newGen(factorGenB, testInputGenB, multiplesGenB)
	assert.Equal(t, 1, judge(genA, genB, 1056))

	genA = newGen(factorGenA, testInputGenA, multiplesGenA)
	genB = newGen(factorGenB, testInputGenB, multiplesGenB)
	assert.Equal(t, 309, judge(genA, genB, 5000000))
}
