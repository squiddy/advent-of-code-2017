package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Program struct {
	name   string
	weight int
	childs []string
}

func parse(filename string) map[string]Program {
	programs := make(map[string]Program)

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ",", "", -1)
		fields := strings.Fields(line)
		name := fields[0]

		var childs []string
		if len(fields) < 3 {
			childs = []string{}
		} else {
			childs = fields[3:]
		}

		programs[name] = Program{
			name:   name,
			weight: 0,
			childs: childs,
		}
	}

	return programs
}

func Index(haystack []string, needle string) int {
	for idx, element := range haystack {
		if element == needle {
			return idx
		}
	}

	return -1
}

func Remove(list []string, element string) []string {
	idx := Index(list, element)
	return append(list[:idx], list[idx+1:]...)
}

func findRoot(programs map[string]Program) string {
	// delete all programs that are childs of another program
	var programNames []string
	for name := range programs {
		programNames = append(programNames, name)
	}

	for _, p := range programs {
		for _, child := range p.childs {
			programNames = Remove(programNames, child)
		}
	}

	return programNames[0]
}

func main() {
	programs := parse("../input2.txt")
	fmt.Println(findRoot(programs))
}
