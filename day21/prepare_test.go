package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	raw := ".#./..#/###"
	b := parse(raw)

	assert.Equal(t, 3, b.size)
	assert.Equal(t, strconv.FormatInt(2+32+64+128+256, 16), b.key(0, 0, b.size))
	assert.Equal(t, raw, b.String())
}

func TestAllTransform(t *testing.T) {
	raw := ".#./#../#.."
	b := parse(raw)
	all := allTransforms(b)

	assert.Len(t, all, 8)
}
