package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func walkMaze(instructions []int) int {
	iptr := 0
	step := 0

	for iptr < len(instructions) {
		jump := instructions[iptr]
		instructions[iptr]++
		iptr += jump
		step++
	}

	return step
}

func parseInstruction(r io.Reader) []int {
	var instructions []int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		instructions = append(instructions, value)
	}

	return instructions
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	instructions := parseInstruction(bytes.NewReader(input))

	fmt.Print(walkMaze(instructions))
}
