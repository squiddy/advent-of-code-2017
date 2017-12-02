package main

import (
	"strings"
	"testing"

	"github.com/stvp/assert"
)

func TestParseGrid(t *testing.T) {
	input := `
		1	4  3	 9
		10	-120 12	1`

	assert.Equal(t, parseGrid(strings.NewReader(input)),
		[][]int{
			[]int{1, 4, 3, 9},
			[]int{10, -120, 12, 1},
		})
}

func TestChecksum(t *testing.T) {
	grid := [][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8},
	}
	assert.Equal(t, checksum(grid), 18)
}

func TestChecksum2(t *testing.T) {
	grid := [][]int{
		[]int{5, 9, 2, 8},
		[]int{9, 4, 7, 3},
		[]int{3, 8, 6, 5},
	}
	assert.Equal(t, checksum2(grid), 9)
}
