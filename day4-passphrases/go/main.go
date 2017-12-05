package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkPassword(password string) bool {
	words := strings.Fields(password)

	seen := make(map[string]int)
	for _, word := range words {
		_, exists := seen[word]
		if exists {
			return false
		}
		seen[word] = 1
	}

	return true
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := scanner.Text()
		if checkPassword(password) {
			valid++
		}
	}
	fmt.Print(valid)
}
