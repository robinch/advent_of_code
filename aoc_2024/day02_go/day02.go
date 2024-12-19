package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	reports := parseInput("input")
	nrOfSafeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			nrOfSafeReports++
		}
	}

	fmt.Printf("Day 2 Part 1: %d\n", nrOfSafeReports)
}

func Part2() {
	reports := parseInput("input")
	nrOfSafeReports := 0

	// Not happy with trying every solution
	// Might exist a dynamic programming solution that's faster
	for _, report := range reports {
		if isSafe(report) {
			nrOfSafeReports++
		} else {
			for i, _ := range report {
				reportCopy := slices.Clone(report)
				reportCopy = slices.Delete(reportCopy, i, i+1)

				if isSafe(reportCopy) {
					nrOfSafeReports++
					break
				}
			}
		}
	}

	fmt.Printf("Day 2 Part 2: %d\n", nrOfSafeReports)
}

func isSafe(report []int) bool {
	asc := report[0] < report[1]
	var min, max int

	if asc {
		min = 1
		max = 3
	} else {
		min = -3
		max = -1
	}

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff < min || diff > max {
			return false
		}
	}

	return true
}

func parseInput(path string) [][]int {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := make([][]int, 0)

	for scanner.Scan() {
		row := scanner.Text()
		rowNumbers := make([]int, 0)

		for _, numberAsStrig := range strings.Split(row, " ") {
			number, err := strconv.Atoi(numberAsStrig)

			if err != nil {
				panic(err)
			}

			rowNumbers = append(rowNumbers, number)
		}

		input = append(input, rowNumbers)
	}

	return input
}

func main() {
	Part1()
	Part2()
}
