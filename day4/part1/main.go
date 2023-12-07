package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	string := string(data)
	lines := strings.Split(string, "\n")

	var winningNumbers [][]int
	var guessNumbers [][]int

	for i, line := range lines {
		cardNumberAndNumbers := strings.Split(line, ":")
		if len(cardNumberAndNumbers) == 2 {
			winningNumbers = append(winningNumbers, make([]int, 0))
			guessNumbers = append(guessNumbers, make([]int, 0))

			numbers := cardNumberAndNumbers[1]
			winningAndGuessLine := strings.Split(numbers, "|")
			winningLine := strings.Split(winningAndGuessLine[0], " ")
			guessLine := strings.Split(winningAndGuessLine[1], " ")
			for _, num := range winningLine {
				num, err := strconv.Atoi(num)
				if err == nil {
					winningNumbers[i] = append(winningNumbers[i], num)
				}
			}
			for _, num := range guessLine {
				num, err := strconv.Atoi(num)
				if err == nil {
					guessNumbers[i] = append(guessNumbers[i], num)
				}
			}
		}
	}

	sum := 0

	for i, winningNumberLine := range winningNumbers {
		guessNumbersLine := guessNumbers[i]
		subSum := 0

		for _, guessNumber := range guessNumbersLine {
			for _, winningNumber := range winningNumberLine {
				if guessNumber == winningNumber {
					if subSum == 0 {
						subSum = 1
					} else {
						subSum *= 2
					}

				}
			}
		}
		sum += subSum

	}

	fmt.Println(sum)
}
