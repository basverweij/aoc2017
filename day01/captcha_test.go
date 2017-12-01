package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptcha(t *testing.T) {
	assert.Equal(t, 3, captcha("1122"))
	assert.Equal(t, 4, captcha("1111"))
	assert.Equal(t, 0, captcha("1234"))
	assert.Equal(t, 9, captcha("91212129"))
}
