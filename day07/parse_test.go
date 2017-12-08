package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseProgramShouldParseValidInput(t *testing.T) {
	p, err := parseProgram("a (1)")

	assert.Nil(t, err)
	assert.Equal(t, &program{"a", 1, nil}, p)
}

func TestParseProgramShouldIncludeSubPrograms(t *testing.T) {
	p, err := parseProgram("a (1) -> b, c, d")

	assert.Nil(t, err)
	assert.Equal(t, &program{"a", 1, []string{"b", "c", "d"}}, p)
}

func TestParseProgramShouldReturnErrorOnInvalidInput(t *testing.T) {
	p, err := parseProgram("foo")

	assert.Equal(t, errInvalidProgramString, err)
	assert.Nil(t, p)
}

func TestParseShouldFindAllPrograms(t *testing.T) {
	p, err := parse(bytes.NewBufferString("a (1)\na (1) -> b, c, d\n"))

	assert.Nil(t, err)
	assert.Equal(t, &program{"a", 1, nil}, p[0])
	assert.Equal(t, &program{"a", 1, []string{"b", "c", "d"}}, p[1])
}

func TestParseShouldReturnErrorOnInvalidInput(t *testing.T) {
	p, err := parse(bytes.NewBufferString("foo"))

	assert.Equal(t, errInvalidProgramString, err)
	assert.Nil(t, p)
}
