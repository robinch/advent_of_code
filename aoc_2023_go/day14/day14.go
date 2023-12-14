package day14

import (
	"bufio"
	"fmt"
	"os"
)

type puzzleInput [][]byte

func (p puzzleInput) String() string {
	str := ""

	for _, line := range p {
		str += fmt.Sprintf("%s\n", string(line))
	}

	return str
}

func Part1(filePath string) int {
	input := parseInput(filePath)
	fmt.Println(input)

	return calculateLoad(input)
}

func calculateLoad(input puzzleInput) int {
	load := 0
	nrOfRows := len(input)
	rowLoad := make([]int, len(input[0]))

	for i := 0; i < len(rowLoad); i++ {
		rowLoad[i] = nrOfRows
	}

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'O' {
				load += rowLoad[col]
				rowLoad[col]--
			}

			if input[row][col] == '#' {
				rowLoad[col] = nrOfRows - (row + 1)
			}
		}
	}

	return load
}

func parseInput(filePath string) puzzleInput {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := puzzleInput{}

	for scanner.Scan() {
		row := []byte(scanner.Text())
		input = append(input, row)
	}

	return input
}
