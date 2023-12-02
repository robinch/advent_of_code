package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const MAX_RED_CUBES = 12
const MAX_GREEN_CUBES = 13
const MAX_BLUE_CUBES = 14

type game struct {
	nr         int
	isPossible bool
}

func Part1() int {
	sum := 0

	file, err := os.Open("inputs/day2.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		game := getGame(scanner.Text())

		if game.isPossible {
			sum += game.nr
		}
	}

	return sum
}

func getGame(line string) game {
	isPossible := true

	gameRegex := `Game ([0-9]+):`
	gameNr := getMaxIntFromRegex(line, gameRegex)

	redRegex := `([0-9]+) red`
	reds := getMaxIntFromRegex(line, redRegex)

	greenRegex := `([0-9]+) green`
	greens := getMaxIntFromRegex(line, greenRegex)

	bluehRegex := `([0-9]+) blue`
	blues := getMaxIntFromRegex(line, bluehRegex)

	if (reds > MAX_RED_CUBES || greens > MAX_GREEN_CUBES || blues > MAX_BLUE_CUBES) {
		isPossible = false
	}

	return game{gameNr, isPossible}
}

func getMaxIntFromRegex(line string, regexPattern string) int {
	intMatches := []int{}

	re := regexp.MustCompile(regexPattern)
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		n, err := strconv.Atoi(match[1])

		if err != nil {
			panic(err)
		}

		intMatches = append(intMatches, n)
	}

	return maxVal(intMatches)	
}


func maxVal(list []int) int {
	max := 0
	for _, n := range list {
		if n > max {
			max = n
		}
	}

	return max
}
