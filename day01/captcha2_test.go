package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptcha2(t *testing.T) {
	assert.Equal(t, 6, captcha2("1212"))
	assert.Equal(t, 0, captcha2("1221"))
	assert.Equal(t, 4, captcha2("123425"))
	assert.Equal(t, 12, captcha2("123123"))
	assert.Equal(t, 4, captcha2("12131415"))
}
