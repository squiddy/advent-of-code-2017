package main

import (
	"testing"

	"github.com/stvp/assert"
)

func TestHighestBank(t *testing.T) {
	assert.Equal(t, highestBank([]int{0, 1, 7, 2, 4}), 2)
	assert.Equal(t, highestBank([]int{1, 1, 0, 0, 0}), 0)
	assert.Equal(t, highestBank([]int{4, 1, 4, 2, 8}), 4)
}

func TestRedistribute(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	redistribute(banks)

	assert.Equal(t, banks, []int{2, 4, 1, 2})
}

func TestRunCycles(t *testing.T) {
	assert.Equal(t, runCycles([]int{0, 2, 7, 0}), 5)
}
