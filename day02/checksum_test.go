package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `5	1	9	5
7	5	3
2	4	6	8`

func TestChecksum(t *testing.T) {
	data, err := readData(strings.NewReader(input))
	assert.Nil(t, err)

	assert.Equal(t, 18, checksum(data))
}
