package main

import (
	"testing"

	"github.com/stvp/assert"
)

func TestScore(t *testing.T) {
	assert.Equal(t, score("{}"), 1)
	assert.Equal(t, score("{{{}}}"), 6)
	assert.Equal(t, score("{{},{}}"), 5)
	assert.Equal(t, score("{{{},{},{{}}}}"), 16)
	assert.Equal(t, score("{<a>,<a>,<a>,<a>}"), 1)
	assert.Equal(t, score("{{<ab>},{<ab>},{<ab>},{<ab>}}"), 9)
	assert.Equal(t, score("{{<!!>},{<!!>},{<!!>},{<!!>}}"), 9)
	assert.Equal(t, score("{{<a!>},{<a!>},{<a!>},{<ab>}}"), 3)
}
