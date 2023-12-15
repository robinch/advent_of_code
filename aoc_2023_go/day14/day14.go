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
	tiltNorth(input)
	return calcLoad(input)
}

func Part2(filePath string) int {
	input := parseInput(filePath)
	calc :=calcLoad(input)

	seenVals := []int{}
	lastSeen := map[string]int{}

	n := 0
	cycleBegin := 0

	for true {
		tiltAllDirections(input)
		seenVals = append(seenVals, calc)

		calc = calcLoad(input)

		if i, ok := lastSeen[input.String()]; ok {
			cycleBegin = i
			break
		}

		lastSeen[input.String()] = n
		n++
	}

	cycles := 1000000000
	offset := cycleBegin - 1
	cycle := n - cycleBegin

	index := (cycles - offset) % cycle

	return seenVals[offset + index]
}

func tiltAllDirections(input puzzleInput) {
	tiltNorth(input)
	tiltWest(input)
	tiltSouth(input)
	tiltEast(input)
}

func tiltNorth(input puzzleInput) {
	tiltV(input, false)
}

func tiltSouth(input puzzleInput) {
	tiltV(input, true)
}

func tiltWest(input puzzleInput) {
	tiltH(input, false)
}

func tiltEast(input puzzleInput) {
	tiltH(input, true)
}

func tiltV(input puzzleInput, reverse bool) {
	for col := 0; col < len(input[0]); col++ {
		occurances := 0
		start := 0
		for row := 0; row < len(input); row++ {
			if input[row][col] == 'O' {
				occurances++
			}

			if input[row][col] == '#' {
				shiftCol(input, col, start, row, occurances, reverse)
				occurances = 0
				start = row + 1
			}
		}
		shiftCol(input, col, start, len(input), occurances, reverse)
	}
}

func shiftCol(input puzzleInput, col, startRow, stopRow, occr int, reverse bool) {
	for row := startRow; row < stopRow; row++ {
		if !reverse && row < startRow+occr || reverse && row >= stopRow-occr {
			input[row][col] = 'O'
		} else {
			input[row][col] = '.'
		}
	}
}

func tiltH(input puzzleInput, reverse bool) {
	for row := 0; row < len(input); row++ {
		occurances := 0
		start := 0
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'O' {
				occurances++
			}

			if input[row][col] == '#' {
				shiftRow(input[row][start:col], occurances, reverse)
				occurances = 0
				start = col + 1
			}
		}
		shiftRow(input[row][start:], occurances, reverse)
	}
}

func shiftRow(row []byte, occr int, reverse bool) {
	for i := 0; i < len(row); i++ {
		if !reverse && i < occr || reverse && i >= len(row)-occr {
			row[i] = 'O'
		} else {
			row[i] = '.'
		}
	}
}

func calcLoad(input puzzleInput) int {
	load := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'O' {
				load += len(input) - row
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
