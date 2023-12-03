package day3

import (
	"bufio"
	"os"
	"unicode"
)

func Part2() int {
	sum := 0
	gears := []point{}
	numbers := [][]number{}

	file, err := os.Open("inputs/day3.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		newNumbers, newGears := parseInputPart2(scanner.Text(), y)
		numbers = append(numbers, newNumbers)

		for _, newGear := range newGears {
			gears = append(gears, newGear)
		}
		y++
	}

	for _, gear := range gears {
		adjacentNumbers := getAdjacentNumbers(gear, numbers)

		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0].val * adjacentNumbers[1].val
		}
	}

	return sum
}

func parseInputPart2(line string, y int) ([]number, []point) {
	numbers := []number{}
	gears := []point{}
	lineAsRunes := []rune(line)

	i := 0
	for i < len(lineAsRunes) {
		val := lineAsRunes[i]

		if unicode.IsDigit(val) {
			number := getNumber(lineAsRunes, i, y)
			numbers = append(numbers, number)
			i += number.length

		} else if val == '*' {
			gears = append(gears, point{i, y})
			i++
		} else {
			i++
		}
	}

	return numbers, gears
}

func getAdjacentNumbers(gear point, numbers [][]number) []number {
	adjacentNumbers := []number{}
	y := gear.y
	above := y - 1
	below := y + 1

	if above >= 0 {
		for _, n := range numbers[above] {
			if isAdjacentNumber(gear, n) {
				adjacentNumbers = append(adjacentNumbers, n)
			}
		}
	}

	if below < len(numbers) {
		for _, n := range numbers[below] {
			if isAdjacentNumber(gear, n) {
				adjacentNumbers = append(adjacentNumbers, n)
			}
		}
	}

	for _, n := range numbers[gear.y] {
		if isAdjacentNumber(gear, n) {
			adjacentNumbers = append(adjacentNumbers, n)
		}
	}

	return adjacentNumbers
}

// skip checking y, it's already checked in previous logic
func isAdjacentNumber(gear point, number number) bool {
	return number.x-1 <= gear.x && gear.x <= number.x+number.length
}
