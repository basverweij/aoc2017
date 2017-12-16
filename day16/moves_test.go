package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSpin(t *testing.T) {
	s := newSpin(3)
	assert.NotNil(t, s)
}

func TestSpinString(t *testing.T) {
	s := newSpin(3)
	assert.Equal(t, "s3", s.String())
}

func TestPerformSpin(t *testing.T) {
	s := newSpin(1)
	l := newLine(5)

	s.perform(l)

	assert.Equal(t, []byte{'e', 'a', 'b', 'c', 'd'}, l.d)
}

func TestNewExchange(t *testing.T) {
	e := newExchange(3, 4)
	assert.NotNil(t, e)
}

func TestExchangeString(t *testing.T) {
	e := newExchange(3, 4)
	assert.Equal(t, "x3/4", e.String())
}

func TestPerformExchange(t *testing.T) {
	e := newExchange(3, 4)
	l := newLine(5)
	e.perform(l)

	assert.Equal(t, []byte{'a', 'b', 'c', 'e', 'd'}, l.d)
}

func TestNewPartner(t *testing.T) {
	p := newPartner('b', 'd')
	assert.NotNil(t, p)
}

func TestPartnerString(t *testing.T) {
	p := newPartner('b', 'd')
	assert.Equal(t, "pb/d", p.String())
}

func TestPerformPartner(t *testing.T) {
	p := newPartner('b', 'd')
	l := newLine(5)

	p.perform(l)

	assert.Equal(t, []byte{'a', 'd', 'c', 'b', 'e'}, l.d)
}
