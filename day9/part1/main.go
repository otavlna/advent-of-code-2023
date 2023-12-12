package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	checkErr(err)

	histories := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		if len(values) == 1 {
			continue
		}
		histories = append(histories, nil)
		for _, str := range values {
			history := &histories[len(histories)-1]
			value, err := strconv.Atoi(str)
			checkErr(err)
			*history = append(*history, value)
		}
	}

	// fmt.Println(histories)

	sum := 0
	for _, history := range histories {
		sequences := make([][]int, 0)
		sequences = append(sequences, history)
	out:
		for {
			lastSequence := sequences[len(sequences)-1]
			newSequence := make([]int, 0)
			for i := 1; i < len(lastSequence); i++ {
				newSequence = append(newSequence, lastSequence[i]-lastSequence[i-1])
			}
			sequences = append(sequences, newSequence)
			for _, value := range newSequence {
				if value != 0 {
					continue out
				}
			}
			break
		}

		// fmt.Println(sequences)

		sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)
		for i := len(sequences) - 2; i >= 0; i-- {
			currentSeq := sequences[i]
			lowerSeq := sequences[i+1]
			lastOfCurrentSeq := currentSeq[len(currentSeq)-1]
			lastOfLowerSeq := lowerSeq[len(lowerSeq)-1]
			newValue := lastOfCurrentSeq + lastOfLowerSeq
			sequences[i] = append(sequences[i], newValue)
		}
		sum += sequences[0][len(sequences[0])-1]
	}

	fmt.Println(sum)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
