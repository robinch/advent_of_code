package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result  int
	numbers []int
}

func part1() {
	equations := parseInput("input.txt")

	calibrationResult := 0

	for _, eq := range equations {
		if check(eq.result, eq.numbers[1:], eq.numbers[0]) {
			calibrationResult += eq.result
		}
	}

	fmt.Printf("Day 7, part 1: %d\n", calibrationResult)
}

func check(result int, numbers []int, acc int) bool {
	// Base case
	if len(numbers) == 0 {
		if acc == result {
			return true
		} else {
			return false
		}
	}

	// Recursion
	return check(result, numbers[1:], acc+numbers[0]) ||
		check(result, numbers[1:], acc*numbers[0])
}

func part2() {
	equations := parseInput("input.txt")

	calibrationResult := 0

	for _, eq := range equations {
		if checkWithConcat(eq.result, eq.numbers[1:], eq.numbers[0]) {
			calibrationResult += eq.result
		}
	}

	fmt.Printf("Day 7, part 2: %d\n", calibrationResult)
}

func checkWithConcat(result int, numbers []int, acc int) bool {
	// Base case
	if len(numbers) == 0 {
		if acc == result {
			return true
		} else {
			return false
		}
	}

	// Recurcion
	concatenated, err := strconv.Atoi(strconv.Itoa(acc) + strconv.Itoa(numbers[0]))

	if err != nil {
		panic(err)
	}

	return checkWithConcat(result, numbers[1:], acc+numbers[0]) ||
		checkWithConcat(result, numbers[1:], acc*numbers[0]) ||
		checkWithConcat(result, numbers[1:], concatenated)
}

func parseInput(path string) []equation {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	equations := make([]equation, 0)

	for scanner.Scan() {
		row := scanner.Text()
		splitRow := strings.Split(row, ": ")

		result, err := strconv.Atoi(splitRow[0])

		if err != nil {
			panic(err)
		}

		numbersAsStrings := strings.Split(splitRow[1], " ")

		numbers := make([]int, len(numbersAsStrings))

		for i, n := range numbersAsStrings {
			number, err := strconv.Atoi(n)

			if err != nil {
				panic(err)
			}

			numbers[i] = number
		}

		equations = append(equations, equation{result: result, numbers: numbers})
	}

	return equations
}

func main() {
	part1()
	part2()
}
