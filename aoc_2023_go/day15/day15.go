package day15

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type lense struct {
	label       string
	focalLength int
}

type step struct {
	lense      lense
	opertation byte
}

func Part1(filePath string) int {
	input := parseInput(filePath)
	sum := 0

	for _, s := range input {
		sum += hash(s)
	}

	return sum
}

func Part2(filePath string) int {
	sum := 0
	input := parseInput(filePath)
	steps := toSteps(input)
	hm := newHashMap()

	for _, s := range steps {
		hm.doStep(s)
	}

	for i, l := range hm {
		slot := 1
		if l.Len() > 0 {
			for current := l.head; current != nil; current = current.next {
				sum += (i + 1) * slot * current.val.focalLength
				slot++
			}
		}
	}

	return sum
}

func toSteps(input []string) []step {
	steps := []step{}
	for _, s := range input {
		split := strings.Split(s, "=")
		label := split[0]
		if len(split) == 2 {
			n, err := strconv.Atoi(split[1])

			if err != nil {
				panic(err)
			}

			steps = append(steps, step{lense: lense{label: label, focalLength: n}, opertation: '='})
		} else {
			steps = append(steps, step{lense: lense{label: label[:len(label)-1]}, opertation: '-'})
		}
	}

	return steps
}

func parseInput(filePath string) []string {
	puzzleInputs := []string{}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	puzzleInputs = strings.Split(scanner.Text(), ",")

	return puzzleInputs
}
