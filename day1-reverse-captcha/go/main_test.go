package main

import (
	"testing"

	"github.com/stvp/assert"
)

func TestReverseCaptcha(t *testing.T) {
	assert.Equal(t, reverseCaptcha("1122"), 3)
	assert.Equal(t, reverseCaptcha("1111"), 4)
	assert.Equal(t, reverseCaptcha("1234"), 0)
	assert.Equal(t, reverseCaptcha("91212129"), 9)
}
