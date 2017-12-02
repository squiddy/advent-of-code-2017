package main

import (
	"testing"

	"github.com/stvp/assert"
)

func TestChecksum(t *testing.T) {
	grid := [][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8},
	}
	assert.Equal(t, checksum(grid), 18)
}
