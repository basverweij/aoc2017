package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	assert.Equal(t, "{}", scan("{}").String())
	assert.Equal(t, "{{{}}}", scan("{{{}}}").String())
	assert.Equal(t, "{{},{}}", scan("{{},{}}").String())
	assert.Equal(t, "{{{},{},{{}}}}", scan("{{{},{},{{}}}}").String())
	assert.Equal(t, "{}", scan("{<{},{},{{}}>}").String())
	assert.Equal(t, "{}", scan("{<a>,<a>,<a>,<a>}").String())
	assert.Equal(t, "{{},{},{},{}}", scan("{{<a>},{<a>},{<a>},{<a>}}").String())
	assert.Equal(t, "{{}}", scan("{{<!>},{<!>},{<!>},{<a>}}").String())
}
