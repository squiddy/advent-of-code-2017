package main

import (
	"fmt"
	"io/ioutil"
)

func score(input string) int {
	score := 0
	inGarbage := false
	groupLevel := 0
	idx := 0

	for idx < len(input) {
		char := string(input[idx])
		switch char {
		case "!":
			idx++
			break
		case "{":
			if !inGarbage {
				groupLevel++
			}
			break
		case "}":
			if !inGarbage {
				score += groupLevel
				groupLevel--
			}
			break
		case "<":
			inGarbage = true
			break
		case ">":
			inGarbage = false
			break
		default:
			break
		}
		idx++
	}
	return score
}

func main() {
	bytes, _ := ioutil.ReadFile("../input.txt")
	input := string(bytes)
	fmt.Println(score(input))
}
