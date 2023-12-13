package day02

import (
	"bufio"
	"os"
)

type gamePart2 struct {
	nr, power int
}

func Part2() int {
	sum := 0

	file, err := os.Open("inputs/day2.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		game := getGame2(scanner.Text())

		sum += game.power
	}

	return sum
}

func getGame2(line string) gamePart2 {
	gameRegex := `Game ([0-9]+):`
	gameNr := getMaxIntFromRegex(line, gameRegex)

	redRegex := `([0-9]+) red`
	reds := getMaxIntFromRegex(line, redRegex)

	greenRegex := `([0-9]+) green`
	greens := getMaxIntFromRegex(line, greenRegex)

	bluehRegex := `([0-9]+) blue`
	blues := getMaxIntFromRegex(line, bluehRegex)

	power := reds * greens * blues

	return gamePart2{gameNr, power}
}
