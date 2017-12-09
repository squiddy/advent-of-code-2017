package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	programs := make(map[string][]string)

	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ",", "", -1)
		fields := strings.Fields(line)
		if len(fields) < 3 {
			// ignore programs without childs
			continue
		}
		name := fields[0]
		childs := fields[3:]
		programs[name] = childs
	}

	// delete all programs that are childs of another program
	toDelete := []string{}
	for _, childs := range programs {
		for _, child := range childs {
			toDelete = append(toDelete, child)
		}
	}

	for _, name := range toDelete {
		delete(programs, name)
	}

	fmt.Println(programs)
}
