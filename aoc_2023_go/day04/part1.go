package day04

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func Part1(filePath string) int {
	points := 0
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := getMatches(scanner.Text())

		if matches > 0 {
			points += pow(2, matches-1)
		}
	}

	return points
}

func getMatches(line string) int {
	splits := strings.Split(line, ":")
	splits = strings.Split(splits[1], "|")
	winningPart := getNumbers(splits[0])
	yourPart := getNumbers(splits[1])

	matches := 0

	for _, number := range yourPart {
		if contains(winningPart, number) {
			matches++
		}
	}

	return matches
}

func getNumbers(s string) []string {
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	numbers := re.FindAllString(s, -1)
	return numbers
}

func contains(list []string, elem string) bool {
	for _, val := range list {
		if elem == val {
			return true
		}
	}

	return false
}

func pow(base int, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}

	return res
}
