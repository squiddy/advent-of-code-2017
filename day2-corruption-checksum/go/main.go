package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func minMaxDifference(numbers []int) int {
	largest := numbers[0]
	lowest := numbers[0]

	for _, number := range numbers {
		if number < lowest {
			lowest = number
		}
		if number > largest {
			largest = number
		}
	}

	return largest - lowest
}

func checksum(grid [][]int) int {
	checksum := 0

	for _, row := range grid {
		checksum += minMaxDifference(row)
	}

	return checksum
}

func maxDivideable(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			number1 := numbers[i]
			number2 := numbers[j]

			if number2 != 0 && number1%number2 == 0 {
				return number1 / number2
			} else if number1 != 0 && number2%number1 == 0 {
				return number2 / number1
			}
		}
	}

	return 0
}

func checksum2(grid [][]int) int {
	checksum := 0

	for _, row := range grid {
		checksum += maxDivideable(row)
	}

	return checksum
}

func parseGrid(r io.Reader) [][]int {
	var grid [][]int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var row []int
		for _, field := range strings.Fields(scanner.Text()) {
			value, _ := strconv.Atoi(field)
			row = append(row, value)
		}

		if len(row) > 0 {
			grid = append(grid, row)
		}
	}

	return grid
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	grid := parseGrid(bytes.NewReader(input))
	fmt.Print(checksum2(grid))
}
