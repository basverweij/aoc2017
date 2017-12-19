package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `     |          ` + "\r\n" +
	`     |  +--+    ` + "\r\n" +
	`     A  |  C    ` + "\r\n" +
	` F---|----E|--+ ` + "\r\n" +
	`     |  |  |  D ` + "\r\n" +
	`     +B-+  +--+ ` + "\r\n" +
	`                `

func createTubes() [][]byte {
	return parse(testInput)
}

func TestFollow(t *testing.T) {
	tubes := createTubes()
	s, n := follow(tubes)
	assert.Equal(t, "ABCDEF", s)
	assert.Equal(t, 38, n)
}

func TestStart(t *testing.T) {
	tubes := createTubes()
	assert.Equal(t, &position{5, 0}, findStart(tubes))
}
