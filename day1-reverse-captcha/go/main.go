package main

import (
	"fmt"
	"os"
)

func ctoi(char byte) int {
	return int(char - '0')
}

func reverseCaptcha(input string) int {
	sum := 0
	for idx := 0; idx < len(input); idx++ {
		number := ctoi(input[idx])
		nextIdx := (idx + 1) % len(input)
		nextNumber := ctoi(input[nextIdx])

		if number == nextNumber {
			sum += number
		}
	}
	return sum
}

func reverseCaptchaStepTwo(input string) int {
	sum := 0
	for idx := 0; idx < len(input); idx++ {
		number := ctoi(input[idx])
		nextIdx := (idx + (len(input) / 2)) % len(input)
		nextNumber := ctoi(input[nextIdx])

		if number == nextNumber {
			sum += number
		}
	}
	return sum
}

func main() {
	fmt.Println(reverseCaptchaStepTwo(os.Args[1]))
}
