package day11

import (
	"sort"
)

type galaxy struct {
	x, y int
}

func Part1(filePath string) int {
	galaxyMap := parseInputToGalaxyMap(filePath)
	galaxyMap.Expand()
	galaxies := allGalaxies(galaxyMap)

	sumOfDistances := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sumOfDistances += manhattanDistance(galaxies[i], galaxies[j])
		}
	}

	return sumOfDistances
}

func Part2(filePath string, expansionMultiplier int) int {
	galaxyMap := parseInputToGalaxyMap(filePath)
	galaxies := allGalaxies(galaxyMap)
	rowExpansions := galaxyMap.GetRowExpansions()
	columnExpansions := galaxyMap.GetColumnExpansions()

	sumOfDistances := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sumOfDistances += manhattanDistance(galaxies[i], galaxies[j])

			ys := []int{galaxies[i].y, galaxies[j].y}
			sort.Ints(ys)

			for _, expandedRow := range rowExpansions {
				if ys[0] < expandedRow && expandedRow < ys[1] {
					sumOfDistances += expansionMultiplier - 1
				}
			}

			xs := []int{galaxies[i].x, galaxies[j].x}
			sort.Ints(xs)

			for _, expandedColumn := range columnExpansions {
				if xs[0] < expandedColumn && expandedColumn < xs[1] {
					sumOfDistances += expansionMultiplier - 1
				}
			}
		}
	}

	return sumOfDistances
}

func allGalaxies(galaxyMap galaxyMap) []galaxy {
	galaxies := []galaxy{}

	for y, row := range galaxyMap {
		for x, col := range row {
			if col == '#' {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}

	return galaxies
}

func manhattanDistance(g1, g2 galaxy) int {
	return abs(g1.x-g2.x) + abs(g1.y-g2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
