package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `0
3
0
1
-3`

func TestParse(t *testing.T) {
	r := bytes.NewReader([]byte(input))

	js, err := parse(r)
	assert.Nil(t, err)
	assert.EqualValues(t, jumps{0, 3, 0, 1, -3}, js)
}
