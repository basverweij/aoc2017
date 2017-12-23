package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVirus(t *testing.T) {
	v := newVirus([]position{{0, 0}})
	assert.NotNil(t, v)
}

func TestVirusInfectedCount(t *testing.T) {
	v := newVirus(parse(rawTestInput))

	assert.Equal(t, 2, v.infectedCount())
}

func TestVirusString(t *testing.T) {
	v := newVirus(parse(rawTestInput))

	assert.Equal(t, " . . # \n"+
		" #[.]. \n"+
		" . . . ",
		v.String(3))
}

func TestVirusBurst(t *testing.T) {
	v := newVirus(parse(rawTestInput))

	for i := 0; i < 2; i++ {
		v.burst()
	}

	assert.Equal(t, 0, v.infections)

	assert.Equal(t, " . . . . . . . . . \n"+
		" . . . . . . . . . \n"+
		" . . . . . . . . . \n"+
		" . . .[.]. # . . . \n"+
		" . . . F W . . . . \n"+
		" . . . . . . . . . \n"+
		" . . . . . . . . . \n"+
		" . . . . . . . . . \n"+
		" . . . . . . . . . ",
		v.String(9))

	v = newVirus(parse(rawTestInput))

	for i := 0; i < 10000000; i++ {
		v.burst()
	}

	assert.Equal(t, 2511944, v.infections)
}
