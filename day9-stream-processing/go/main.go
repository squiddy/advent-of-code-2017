package main

import (
	"fmt"
	"io/ioutil"
)

func garbageCount(input string) int {
	garbageCount := 0
	inGarbage := false
	idx := 0

	for idx < len(input) {
		char := string(input[idx])
		if char == "!" {
			idx += 2
			continue
		}

		switch char {
		case "<":
			if inGarbage {
				garbageCount++
			}
			inGarbage = true
			break
		case ">":
			inGarbage = false
			break
		default:
			if inGarbage {
				garbageCount++
			}
		}
		idx++
	}
	return garbageCount
}

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
	fmt.Println(garbageCount(input))
}
