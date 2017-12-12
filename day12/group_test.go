package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = map[int][]int{
	0: []int{2},
	1: []int{1},
	2: []int{0, 3, 4},
	3: []int{2, 4},
	4: []int{2, 3, 6},
	5: []int{6},
	6: []int{4, 5},
}

func TestGroup(t *testing.T) {
	assert.Equal(t, 6, len(group(testInput, 0)))
}

func TestGroupCount(t *testing.T) {
	assert.Equal(t, 2, groupCount(testInput))
}
