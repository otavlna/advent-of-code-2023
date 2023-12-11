package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times := parseLine(scanner.Text())
	scanner.Scan()
	distances := parseLine(scanner.Text())

	// fmt.Println(times)
	// fmt.Println(distances)

	var product int = 1

	for i, distance := range distances {
		var sum int = 0
		totalTime := times[i]
		for holdTime := 0; holdTime <= totalTime; holdTime++ {
			sum += canBeat(distance, holdTime, totalTime)
		}
		product *= sum
	}

	fmt.Println(product)
}

func canBeat(distance int, holdTime int, totalTime int) int {
	if distance < holdTime*(totalTime-holdTime) {
		return 1
	} else {
		return 0
	}
}

func parseLine(line string) []int {
	slice := make([]int, 0, 4)
	f := func(r rune) bool {
		return !unicode.IsNumber(r)
	}

	values := strings.FieldsFunc(strings.ReplaceAll(line, " ", ""), f)
	for _, val := range values {
		num, err := strconv.Atoi(val)
		checkErr(err)
		slice = append(slice, num)
	}
	return slice
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
