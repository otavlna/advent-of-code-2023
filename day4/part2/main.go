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
	var cardCount []int

	for i, line := range lines {
		cardNumberAndNumbers := strings.Split(line, ":")
		if len(cardNumberAndNumbers) == 2 {
			winningNumbers = append(winningNumbers, make([]int, 0))
			guessNumbers = append(guessNumbers, make([]int, 0))
			cardCount = append(cardCount, 1)

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

	total := 0

	for i, winningNumberLine := range winningNumbers {
		for x := 0; x < cardCount[i]; x++ {
			total++
			guessNumbersLine := guessNumbers[i]
			sum := 0

			for _, guessNumber := range guessNumbersLine {
				for _, winningNumber := range winningNumberLine {
					if guessNumber == winningNumber {
						sum++
					}
				}
			}

			for j := 0; j < sum; j++ {
				cardCount[i+j+1]++
			}
		}
	}

	fmt.Println(total)

}
