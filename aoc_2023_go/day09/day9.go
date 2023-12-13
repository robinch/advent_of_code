package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1(filePath string) int {
	inputs := getInputs(filePath)
	nextInSeries := []int{}

	for _, input := range inputs {
		nextInSeries = append(nextInSeries, getNextInSeries(input))
	}

	return sum(nextInSeries)
}

func Part2(filePath string) int {
	inputs := getInputs(filePath)
	previousInSeries := []int{}

	for _, input := range inputs {
		previousInSeries = append(previousInSeries, getPreviousInSeries(input))
	}

	return sum(previousInSeries)
}

func getNextInSeries(input []int) int {
	lastNumbers := []int{}
	diffs := input
	allZero := false

	for !allZero {
		lastNumbers = append(lastNumbers, diffs[len(diffs)-1])
		diffs = getDiffs(diffs)

		allZero = true
		for _, diff := range diffs {
			if diff != 0 {
				allZero = false
			}
		}
	}

	return sum(lastNumbers)
}

func getPreviousInSeries(input []int) int {
	firstNumbers := []int{}
	diffs := input
	allZero := false

	for !allZero {
		firstNumbers = append(firstNumbers, diffs[0])
		diffs = getDiffs(diffs)

		allZero = true
		for _, diff := range diffs {
			if diff != 0 {
				allZero = false
			}
		}
	}

	return subtract(firstNumbers)
}

func getDiffs(input []int) []int {
	diffs := []int{}

	for i := 1; i < len(input); i++ {
		diffs = append(diffs, input[i]-input[i-1])
	}

	return diffs
}

func sum(input []int) int {
	total := 0

	for _, n := range input {
		total += n
	}

	return total
}

func subtract(input []int) int {
	total := 0

	for i := len(input) - 1; i >= 0; i-- {
		total = input[i] - total
	}

	return total
}

func getInputs(filePath string) [][]int {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := [][]int{}

	for scanner.Scan() {
		inputs = append(inputs, getNumbers(scanner.Text()))
	}

	return inputs
}

func getNumbers(input string) []int {
	numbers := []int{}

	for _, s := range strings.Split(input, " ") {
		n, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	return numbers
}
