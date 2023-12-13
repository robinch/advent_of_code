package day04

import (
	"bufio"
	"os"
)

func Part2(filePath string) int {
	scratchCards := 0
	file, err := os.Open(filePath)
	cardCopiesWon := [200]int{}

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	card := 1
	for scanner.Scan() {
		wins := getMatches(scanner.Text())
		copiesOfCard := 1 + cardCopiesWon[card]

		scratchCards += copiesOfCard

		for i := 0; i < wins; i++ {
			cardCopiesWon[card+i+1] += copiesOfCard
		}

		card++
	}

	return scratchCards
}
