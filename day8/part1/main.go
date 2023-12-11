package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type NodeValues struct {
	L string
	R string
}

func main() {
	file, err := os.Open("../input.txt")
	checkErr(err)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructionsLine := scanner.Text()
	instructions := make([]rune, len(instructionsLine))
	for i, r := range instructionsLine {
		instructions[i] = r
	}

	// fmt.Printf("%c\n", instructions)

	nodes := make(map[string]NodeValues)

	for scanner.Scan() {
		line := scanner.Text()
		f := func(c rune) bool {
			return !unicode.IsLetter(c)
		}
		values := strings.FieldsFunc(line, f)
		if len(values) != 3 {
			continue
		}
		nodes[values[0]] = NodeValues{values[1], values[2]}
	}

	// fmt.Println(nodes)

	steps := 0
	current := "AAA"
	for i := 0; ; i = (i + 1) % len(instructions) {
		if instructions[i] == 'L' {
			current = nodes[current].L
		} else {
			current = nodes[current].R
		}
		steps++
		if current == "ZZZ" {
			break
		}
	}

	fmt.Println(steps)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
