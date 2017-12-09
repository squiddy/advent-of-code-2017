package main

import (
	"fmt"
)

func highestBank(banks []int) int {
	var highestIdx int
	var highestValue int

	for idx, value := range banks {
		if value > highestValue {
			highestIdx = idx
			highestValue = value
		}
	}

	return highestIdx
}

func redistribute(banks []int) {
	idx := highestBank(banks)

	distribute := banks[idx]
	banks[idx] = 0

	for distribute > 0 {
		idx++
		idx = idx % len(banks)
		distribute--
		banks[idx]++
	}
}

func compare(a, b []int) bool {
	for idx, value := range a {
		if b[idx] != value {
			return false
		}
	}
	return true
}

func contains(list [][]int, needle []int) bool {
	for _, d := range list {
		if compare(d, needle) {
			return true
		}
	}

	return false
}

func runCycles(banks []int) int {
	cycle := 0
	distributions := [][]int{banks}

	currentBank := make([]int, len(banks))
	copy(currentBank, banks)

	for {
		redistribute(currentBank)
		cycle++

		if contains(distributions, currentBank) {
			break
		}

		c := make([]int, len(currentBank))
		copy(c, currentBank)
		distributions = append(distributions, c)
	}

	return cycle
}

func main() {
	banks := []int{
		2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14,
	}

	fmt.Println(runCycles(banks))
}
