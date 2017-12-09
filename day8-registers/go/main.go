package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Lines(filename string) []string {
	lines := []string{}

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

type Instruction struct {
	register        string
	instruction     string
	value           int
	compareOp       string
	compareRegister string
	compareOperand  int
}

func ParseInstruction(line string) Instruction {
	fields := strings.Fields(line)

	parseInt := func(text string) int {
		value, _ := strconv.Atoi(text)
		return value
	}

	return Instruction{
		register:        fields[0],
		instruction:     fields[1],
		value:           parseInt(fields[2]),
		compareRegister: fields[4],
		compareOp:       fields[5],
		compareOperand:  parseInt(fields[6]),
	}
}

func CreateProgram(filename string) []Instruction {
	var instructions []Instruction

	for _, line := range Lines(filename) {
		instr := ParseInstruction(line)
		instructions = append(instructions, instr)
	}

	return instructions
}

func RunCompare(instr Instruction, registers map[string]int) bool {
	registerValue := registers[instr.compareRegister]
	switch instr.compareOp {
	case "<":
		return registerValue < instr.compareOperand
	case "<=":
		return registerValue <= instr.compareOperand
	case ">":
		return registerValue > instr.compareOperand
	case ">=":
		return registerValue >= instr.compareOperand
	case "==":
		return registerValue == instr.compareOperand
	case "!=":
		return registerValue != instr.compareOperand
	}
	panic("Unrecognized comparision")
}

func RunProgram(program []Instruction) map[string]int {
	registers := make(map[string]int)

	for _, instr := range program {
		if !RunCompare(instr, registers) {
			continue
		}

		_, present := registers[instr.register]
		if !present {
			registers[instr.register] = 0
		}

		switch instr.instruction {
		case "inc":
			registers[instr.register] += instr.value
		case "dec":
			registers[instr.register] -= instr.value
		default:
			panic("Unrecognized instruction")
		}
	}

	return registers
}

func main() {
	program := CreateProgram("../input2.txt")
	registers := RunProgram(program)

	highest := 0
	for _, value := range registers {
		if value > highest {
			highest = value
		}
	}
	fmt.Println(highest)
}
