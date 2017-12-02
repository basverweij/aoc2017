package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input2 = `5	9	2	8
9	4	7	3
3	8	6	5`

func TestChecksum2(t *testing.T) {
	data, err := readData(strings.NewReader(input2))
	assert.Nil(t, err)

	assert.Equal(t, 9, checksum2(data))
}
