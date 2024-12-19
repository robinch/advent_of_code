package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Part1() {
	input := parseInput("input")

	searchedWord := []rune("XMAS")

	rowMatches := rowMatches(input, searchedWord)
	colMatches := colMatches(input, searchedWord)
	rightDiagonalMatches := rightDiagonalMatches(input, searchedWord)
	leftDiagonalMatches := leftDiagonalMatches(input, searchedWord)

	fmt.Printf("Day 4 Part 1: %v\n", rowMatches+colMatches+rightDiagonalMatches+leftDiagonalMatches)
}

func Part2() {
	matches := 0
	input := parseInput("input")
	searchedWord := []rune("MAS")

	searchedWordReverse := slices.Clone(searchedWord)
	slices.Reverse(searchedWordReverse)
	wordLength := len(searchedWord)

	for row := 0; row < len(input)-wordLength+1; row++ {
		for col := 0; col < len(input[0])-wordLength+1; col++ {
			rightDiagnalWord := make([]rune, wordLength)
			leftDiagnalWord := make([]rune, wordLength)
			for i := 0; i < wordLength; i++ {
				rightDiagnalWord[i] = input[row+i][col+i]
				leftDiagnalWord[i] = input[row+i][col+wordLength-1-i]
			}

			if (slices.Equal(rightDiagnalWord, searchedWord) || slices.Equal(rightDiagnalWord, searchedWordReverse)) &&
				(slices.Equal(leftDiagnalWord, searchedWord) || slices.Equal(leftDiagnalWord, searchedWordReverse)) {
				matches++
			}
		}
	}

	fmt.Printf("Day 4 Part 2: %v\n", matches)
}

func rowMatches(input [][]rune, searchedWord []rune) int {
	matches := 0
	searchedWordReverse := slices.Clone(searchedWord)
	slices.Reverse(searchedWordReverse)
	wordLength := len(searchedWord)

	for _, row := range input {
		for i := 0; i < len(row)-wordLength+1; i++ {
			word := row[i : i+wordLength]
			if slices.Equal(word, searchedWord) || slices.Equal(word, searchedWordReverse) {
				matches++
			}

		}
	}

	return matches
}

func colMatches(input [][]rune, searchedWord []rune) int {
	matches := 0
	searchedWordReverse := slices.Clone(searchedWord)
	slices.Reverse(searchedWordReverse)
	wordLength := len(searchedWord)

	for col := 0; col < len(input[0]); col++ {
		for row := 0; row < len(input)-wordLength+1; row++ {
			word := make([]rune, wordLength)
			for i := 0; i < wordLength; i++ {
				word[i] = input[row+i][col]
			}

			if slices.Equal(word, searchedWord) || slices.Equal(word, searchedWordReverse) {
				matches++
			}
		}
	}

	return matches
}

func rightDiagonalMatches(input [][]rune, searchedWord []rune) int {
	matches := 0
	searchedWordReverse := slices.Clone(searchedWord)
	slices.Reverse(searchedWordReverse)
	wordLength := len(searchedWord)

	for row := 0; row < len(input)-wordLength+1; row++ {
		for col := 0; col < len(input[0])-wordLength+1; col++ {
			word := make([]rune, wordLength)
			for i := 0; i < wordLength; i++ {
				word[i] = input[row+i][col+i]
			}

			if slices.Equal(word, searchedWord) || slices.Equal(word, searchedWordReverse) {
				matches++
			}
		}
	}

	return matches
}

func leftDiagonalMatches(input [][]rune, searchedWord []rune) int {
	matches := 0
	searchedWordReverse := slices.Clone(searchedWord)
	slices.Reverse(searchedWordReverse)
	wordLength := len(searchedWord)

	for row := 0; row < len(input)-wordLength+1; row++ {
		for col := wordLength - 1; col < len(input[0]); col++ {

			word := make([]rune, wordLength)
			for i := 0; i < 4; i++ {
				word[i] = input[row+i][col-i]
			}

			if slices.Equal(word, searchedWord) || slices.Equal(word, searchedWordReverse) {
				matches++
			}
		}
	}

	return matches
}

func parseInput(filename string) [][]rune {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, []rune(line))
	}

	return input
}

func main() {
	Part1()
	Part2()
}
