package day06

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
)

func Part1(filePath string) int {
	nrOfWaysToWin := 1
	times := []int{}
	distances := []int{}

	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times = getNumbersFromInput(scanner.Text())
	scanner.Scan()
	distances = getNumbersFromInput(scanner.Text())

	for i := 0; i < len(times); i++ {
		possibleWins := getPossibleWins(distances[i], times[i])

		nrOfWaysToWin *= possibleWins
	}

	return nrOfWaysToWin
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := getNumberFromInput(scanner.Text())
	scanner.Scan()
	distance := getNumberFromInput(scanner.Text())

	return getPossibleWins(distance, time)
}

func getPossibleWins(distance, raceTime int) int {
	recordHold := getHoldTime(distance, raceTime)
	maxDistanceHold := getMaxDistanceHoldTime(raceTime)

	isWholeNumber := false

	if math.Trunc(maxDistanceHold) == maxDistanceHold {
		isWholeNumber = true
	}

	possibleWins := 2 * (int(math.Trunc(maxDistanceHold)) - int(math.Trunc(recordHold)))

	if isWholeNumber {
		possibleWins--
	}

	return possibleWins
}

func getHoldTime(distance, raceTime int) float64 {
	d := float64(distance)
	rt := float64(raceTime)
	return (rt - math.Sqrt(math.Pow(rt, 2)-4*d)) / 2
}

func getMaxDistanceHoldTime(raceTime int) float64 {
	return float64(raceTime) / 2
}

func getNumberFromInput(line string) int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)

	s := ""
	for _, match := range matches {
		s += match
	}

	number, err := strconv.Atoi(s)
	check(err)

	return number
}

func getNumbersFromInput(line string) []int {
	numbers := []int{}
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)

	for _, match := range matches {
		n, err := strconv.Atoi(match)
		check(err)

		numbers = append(numbers, n)
	}

	return numbers
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
