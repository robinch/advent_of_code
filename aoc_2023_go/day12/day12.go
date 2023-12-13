package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type puzzleInput struct {
	springs []byte
	numbers []int
}

func (p puzzleInput) String() string {
	return fmt.Sprintf("{%s, %v}", string(p.springs), p.numbers)
}

func Part1(filePath string) int {
	sum := 0
	puzzleInputs := parseInput(filePath)

	cache := map[string]int{}
	for _, puzzleInput := range puzzleInputs {
		sum += nrOfConfigurations(puzzleInput.springs, puzzleInput.numbers, cache)
	}

	return sum
}

func Part2(filePath string) int {
	sum := 0
	puzzleInputs := parseInput(filePath)
	unfoldedPuzzleInputs := unfold(puzzleInputs)

	cache := map[string]int{}

	for _, puzzleInput := range unfoldedPuzzleInputs {
		sum += nrOfConfigurations(puzzleInput.springs, puzzleInput.numbers, cache)
	}

	return sum
}

func unfold(puzzleInputs []puzzleInput) []puzzleInput {
	unfoldedPuzzleInputs := []puzzleInput{}

	for _, puzzleInput := range puzzleInputs {
		springExpansion := append([]byte{'?'}, puzzleInput.springs...)
		numbersExpansion := []int(puzzleInput.numbers)

		for i := 0; i < 4; i++ {
			puzzleInput.springs = append(puzzleInput.springs, springExpansion...)
			puzzleInput.numbers = append(puzzleInput.numbers, numbersExpansion...)
		}

		unfoldedPuzzleInputs = append(unfoldedPuzzleInputs, puzzleInput)
	}
	return unfoldedPuzzleInputs
}

func nrOfConfigurations(springs []byte, numbers []int, cache map[string]int) int {
	key := puzzleInput{springs, numbers}.String()
	if val, ok := cache[key]; ok {
		return val
	}

	nrOfConfigs := 0

	if len(springs) == 0 {
		if len(numbers) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(numbers) == 0 {
		if !contains('#', springs) {
			return 1
		} else {
			return 0
		}
	}

	if springs[0] == '.' || springs[0] == '?' {
		nrOfConfigs += nrOfConfigurations(springs[1:], numbers, cache)
	}

	if springs[0] == '#' || springs[0] == '?' {
		if numbers[0] == len(springs) && !contains('.', springs[:numbers[0]]) {
			nrOfConfigs += nrOfConfigurations(springs[numbers[0]:], numbers[1:], cache)
		}

		if numbers[0] < len(springs) &&
			!contains('.', springs[:numbers[0]]) &&
			springs[numbers[0]] != '#' {

			nrOfConfigs += nrOfConfigurations(springs[numbers[0]+1:], numbers[1:], cache)
		}
	}

	cache[key] = nrOfConfigs

	return nrOfConfigs
}

func contains(elem byte, elems []byte) bool {
	for _, e := range elems {
		if e == elem {
			return true
		}
	}

	return false
}

func parseInput(filePath string) []puzzleInput {
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	puzzleInputs := []puzzleInput{}

	for scanner.Scan() {
		springs, numbers := parseInputRow(scanner.Text())
		puzzleInputs = append(puzzleInputs, puzzleInput{springs, numbers})
	}

	return puzzleInputs
}

func parseInputRow(row string) ([]byte, []int) {
	split := strings.Split(row, " ")
	springs := []byte(split[0])
	numbers := []int{}

	for _, s := range strings.Split(split[1], ",") {
		num, err := strconv.Atoi(s)
		check(err)
		numbers = append(numbers, num)
	}

	return springs, numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
