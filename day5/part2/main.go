package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type MapLine struct {
	source      uint
	destination uint
	length      uint
}

func getMapOutput(m *[]MapLine, input uint) uint {
	for _, mapLine := range *m {
		if input >= mapLine.source && input < mapLine.source+mapLine.length {
			return mapLine.destination + (input - mapLine.source)
		}
	}
	return input
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var seeds []uint = make([]uint, 0)
	var maps [][]MapLine = make([][]MapLine, 0, 7)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(seeds) == 0 {
			for _, seed := range strings.Split(strings.Split(line, ": ")[1], " ") {
				seedInt, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, uint(seedInt))
				continue
			}
		}

		if line == "" {
			maps = append(maps, make([]MapLine, 0, 50))
			continue
		}

		r, _ := utf8.DecodeRuneInString(line)

		if unicode.IsDigit(r) {
			numbers := strings.Split(line, " ")

			destination, err1 := strconv.Atoi(numbers[0])
			source, err2 := strconv.Atoi(numbers[1])
			length, err3 := strconv.Atoi(numbers[2])

			checkErr(err1)
			checkErr(err2)
			checkErr(err3)

			maps[len(maps)-1] = append(maps[len(maps)-1], MapLine{uint(source), uint(destination), uint(length)})
		}
	}

	//	fmt.Println(seeds)
	//	fmt.Println(maps)

	var lowest uint = math.MaxUint
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedLength := seeds[i+1]
		for seed := seedStart; seed < seedStart+seedLength; seed++ {
			//fmt.Println(seed)
			input := seed
			for _, m := range maps {
				input = getMapOutput(&m, input)
			}
			if input < lowest {
				fmt.Println("new lowest ", input)
				lowest = input
			}
		}
	}

	fmt.Println(lowest)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
