package day01

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type LetterMatch struct {
	offsetFromStart, offsetFromEnd, val int
}

func Part2() int {
	sum := 0

	file, err := os.Open("inputs/day1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sum += getNumberIncludingLetters(scanner.Text())
	}
	return sum
}

func getNumberIncludingLetters(line string) int {
	matches := []LetterMatch{
		getLetterMatch(line, "0", 0),
		getLetterMatch(line, "1", 1),
		getLetterMatch(line, "2", 2),
		getLetterMatch(line, "3", 3),
		getLetterMatch(line, "4", 4),
		getLetterMatch(line, "5", 5),
		getLetterMatch(line, "6", 6),
		getLetterMatch(line, "7", 7),
		getLetterMatch(line, "8", 8),
		getLetterMatch(line, "9", 9),
		getLetterMatch(line, "zero", 0),
		getLetterMatch(line, "one", 1),
		getLetterMatch(line, "two", 2),
		getLetterMatch(line, "three", 3),
		getLetterMatch(line, "four", 4),
		getLetterMatch(line, "five", 5),
		getLetterMatch(line, "six", 6),
		getLetterMatch(line, "seven", 7),
		getLetterMatch(line, "eight", 8),
		getLetterMatch(line, "nine", 9),
	}

	sort.Slice(matches, func(i, j int) bool { return matches[i].offsetFromStart < matches[j].offsetFromStart })

	firstNumber := 10 * matches[0].val

	sort.Slice(matches, func(i, j int) bool { return matches[i].offsetFromEnd < matches[j].offsetFromEnd })

	secondNumber := matches[0].val

	return firstNumber + secondNumber
}

func getLetterMatch(line string, match string, val int) LetterMatch {
	split := strings.Split(line, match)
	offsetFromStart := len(split[0])
	offsetFromEnd := len(split[len(split)-1])

	return LetterMatch{offsetFromStart, offsetFromEnd, val}
}
