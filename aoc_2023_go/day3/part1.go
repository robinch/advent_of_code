package day3

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

type number struct {
	x, y, length, val int
}

type point struct {
	x, y int
}

func Part1() int {
	sum := 0
	hasSymbol := make(map[point]bool)
	numbers := []number{}

	file, err := os.Open("inputs/day3.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		for _, newNumber := range parseInput(scanner.Text(), y, hasSymbol) {
			numbers = append(numbers, newNumber)
		}
		y++
	}

	for _, number := range numbers {
		if hasAdjacentSymbol(number, hasSymbol) {
			sum += number.val
		}
	}

	return sum
}

func parseInput(line string, y int, hasSymbol map[point]bool) []number {
	numbers := []number{}
	lineAsRunes := []rune(line)

	i := 0
	for i < len(lineAsRunes) {
		val := lineAsRunes[i]

		if unicode.IsDigit(val) {
			number := getNumber(lineAsRunes, i, y)
			numbers = append(numbers, number)
			i += number.length

		} else if val != '.' {
			hasSymbol[point{i, y}] = true
			i++
		} else {
			i++
		}
	}

	return numbers
}

func getNumber(lineAsRunes []rune, startIndex, y int) number {
	numberFound := []rune{}

	for i := startIndex; i < len(lineAsRunes) && unicode.IsDigit(lineAsRunes[i]); i++ {
		numberFound = append(numberFound, lineAsRunes[i])
	}

	n, err := strconv.Atoi(string(numberFound))
	if err != nil {
		panic(err)
	}

	return number{startIndex, y, len(numberFound), n}
}

func hasAdjacentSymbol(n number, hasSymbol map[point]bool) bool {
	left := n.x - 1
	right := n.x + n.length
	above := n.y - 1
	under := n.y + 1

	for x := left; x <= right; x++ {
		if hasSymbol[point{x, above}] {
			return true
		}
	}

	for x := left; x <= right; x++ {
		if hasSymbol[point{x, under}] {
			return true
		}
	}

	if hasSymbol[point{left, n.y}] {
		return true
	}

	if hasSymbol[point{right, n.y}] {
		return true
	}

	return false
}
