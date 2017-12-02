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

func TestReverseCaptchaStepTwo(t *testing.T) {
	assert.Equal(t, reverseCaptchaStepTwo("1212"), 6)
	assert.Equal(t, reverseCaptchaStepTwo("1221"), 0)
	assert.Equal(t, reverseCaptchaStepTwo("123425"), 4)
	assert.Equal(t, reverseCaptchaStepTwo("123123"), 12)
	assert.Equal(t, reverseCaptchaStepTwo("12131415"), 4)
}
