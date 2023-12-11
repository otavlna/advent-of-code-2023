package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
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
	currentPositions := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsDigit(c)
		}
		values := strings.FieldsFunc(line, f)
		if len(values) != 3 {
			continue
		}
		nodes[values[0]] = NodeValues{values[1], values[2]}
		if values[0][2] == 'A' {
			currentPositions = append(currentPositions, values[0])
		}
	}

	// fmt.Println(nodes)
	// fmt.Println(currentPositions)

	step := 0
	stepsToFinish := make(map[int]int)
	finishedCount := 0

	for i := 0; ; i = (i + 1) % len(instructions) {
		for j, position := range currentPositions {
			if instructions[i] == 'L' {
				currentPositions[j] = nodes[position].L
			} else {
				currentPositions[j] = nodes[position].R
			}
		}

		step++

		for j, position := range currentPositions {
			if position[2] == 'Z' {
				if stepsToFinish[j] == 0 {
					stepsToFinish[j] = step
					finishedCount++
					// fmt.Println("finished", j, "at instruction", i, "took", step)

				}
			}
		}

		if finishedCount == len(currentPositions) {
			break
		}
	}

	values := maps.Values(stepsToFinish)
	res := LCM(values[0], values[1], values[2:]...)
	fmt.Println(res)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
