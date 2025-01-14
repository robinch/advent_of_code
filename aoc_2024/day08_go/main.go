package main

import (
	"bufio"
	"fmt"
	"os"
)

type location struct {
	x int
	y int
}

type Antennas struct {
	width     int
	height    int
	locations map[rune][]location
}

func part1() {
	antiNodeLocations := make(map[location]struct{})

	antennas := parseInput("input.txt")

	for _, locations := range antennas.locations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diffX := locations[j].x - locations[i].x
				diffY := locations[j].y - locations[i].y

				addIfValid(
					antiNodeLocations,
					location{x: locations[i].x - diffX, y: locations[i].y - diffY},
					antennas.width,
					antennas.height,
				)

				addIfValid(
					antiNodeLocations,
					location{x: locations[j].x + diffX, y: locations[j].y + diffY},
					antennas.width,
					antennas.height,
				)

			}
		}
	}

	fmt.Printf("Day 08, part 1: %v\n", len(antiNodeLocations))
}

func part2() {
	antiNodeLocations := make(map[location]struct{})

	antennas := parseInput("input.txt")

	for _, locations := range antennas.locations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diffX := locations[j].x - locations[i].x
				diffY := locations[j].y - locations[i].y

				for n := 0; addIfValid(antiNodeLocations, location{x: locations[i].x - n*diffX, y: locations[i].y - n*diffY}, antennas.width, antennas.height); n++ {
				}
				for n := 0; addIfValid(antiNodeLocations, location{x: locations[j].x + n*diffX, y: locations[j].y + n*diffY}, antennas.width, antennas.height); n++ {
				}
			}
		}
	}

	fmt.Printf("Day 08, part 2: %v\n", len(antiNodeLocations))
}

func addIfValid(antiNodeLocations map[location]struct{}, locations location, width, height int) bool {
	if locations.x >= 0 && locations.x < width && locations.y >= 0 && locations.y < height {
		antiNodeLocations[locations] = struct{}{}
		return true
	}

	return false
}

func parseInput(path string) Antennas {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	locations := make(map[rune][]location)
	scanner := bufio.NewScanner(file)
	row := 0
	width := -1

	for scanner.Scan() {
		inputRow := scanner.Text()

		if row == 0 {
			width = len(inputRow)
		}

		for col, r := range []rune(inputRow) {
			if r != '.' {
				loc := location{x: col, y: row}

				if val, ok := locations[r]; ok {
					locations[r] = append(val, loc)
				} else {
					locations[r] = []location{loc}
				}
			}
		}

		row++
	}

	return Antennas{width: width, height: row, locations: locations}
}

func main() {
	part1()
	part2()
}
