package main

import (
	"fmt"
	"strconv"
)

func main() {
	word := "ab123uii1j01mmh024zxc01bbb900qq000001"
	count := numbDifferentInteriors(word)
	fmt.Printf("count: %d\n", count)
}

func countDifferentNumbers(arr []int) int {
	numCounts := make(map[int]int)

	for _, num := range arr {
		numCounts[num]++
	}

	return len(numCounts)
}

func numbDifferentInteriors(words string) int {
	numbers := make([]int, 0)
	currentNumberString := ""
	for i := 0; i < len(words); i++ {
		var char = string(words[i])
		_, err := strconv.Atoi(char)

		if err != nil {
			if currentNumberString != "" {
				if finishNumber, err := strconv.Atoi(currentNumberString); err == nil {
					numbers = append(numbers, finishNumber)
					currentNumberString = ""
				}
			}
			continue
		}

		if err == nil {
			currentNumberString = currentNumberString + char
		}

	}
	if currentNumberString != "" {
		if finishNumber, err := strconv.Atoi(currentNumberString); err == nil {
			numbers = append(numbers, finishNumber)
		}
	}

	return countDifferentNumbers(numbers)
}
