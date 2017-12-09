package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func gs(g *group, gc int) string {
	return g.String()
}

func TestScanString(t *testing.T) {
	assert.Equal(t, "{}", gs(scan("{}")))
	assert.Equal(t, "{{{}}}", gs(scan("{{{}}}")))
	assert.Equal(t, "{{},{}}", gs(scan("{{},{}}")))
	assert.Equal(t, "{{{},{},{{}}}}", gs(scan("{{{},{},{{}}}}")))
	assert.Equal(t, "{}", gs(scan("{<{},{},{{}}>}")))
	assert.Equal(t, "{}", gs(scan("{<a>,<a>,<a>,<a>}")))
	assert.Equal(t, "{{},{},{},{}}", gs(scan("{{<a>},{<a>},{<a>},{<a>}}")))
	assert.Equal(t, "{{}}", gs(scan("{{<!>},{<!>},{<!>},{<a>}}")))
}

func gc(g *group, gc int) int {
	return gc
}

func TestScanGarbageCount(t *testing.T) {
	assert.Equal(t, 0, gc(scan("{<>}")))
	assert.Equal(t, 17, gc(scan("{<random characters>}")))
	assert.Equal(t, 3, gc(scan("{<<<<>}")))
	assert.Equal(t, 2, gc(scan("{<{!>}>}")))
	assert.Equal(t, 0, gc(scan("{<!!>}")))
	assert.Equal(t, 0, gc(scan("{<!!!>>}")))
	assert.Equal(t, 10, gc(scan("{<{o\"i!a,<{i<a>}")))
}
