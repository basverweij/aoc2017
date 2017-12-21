package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawTestRules = map[string]string{
	"../.#":       "##./#../...",
	".#./..#/###": "#..#/..../..../#..#",
}

func testInput() (*bitgrid, map[sizedKey]*bitgrid) {
	start := parse(rawStart)
	rules := parseRules(rawTestRules)

	return start, rules
}

func TestInput(t *testing.T) {
	start, rules := testInput()

	assert.NotNil(t, start)
	assert.Len(t, rules, 12)
}
