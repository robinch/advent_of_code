package day11

import (
	"bufio"
	"fmt"
	"os"
)

type galaxyMap [][]byte

func (g galaxyMap) String() string {
	str := ""

	for _, line := range g {
		str += fmt.Sprintf("%s\n", string(line))
	}

	return str
}

func (g galaxyMap) GetRowExpansions() []int {
	expansions := []int{}
	for i := 0; i < len(g); i++ {
		hasGalaxy := false
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '#' {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			expansions = append(expansions, i)
		}
	}
	return expansions
}

func (g galaxyMap) GetColumnExpansions() []int {
	expansions := []int{}
	for i := 0; i < len(g[0]); i++ {
		hasGalaxy := false
		for j := 0; j < len(g); j++ {
			if g[j][i] == '#' {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			expansions = append(expansions, i)
		}
	}

	return expansions
}

func (s *galaxyMap) Expand() {
	*s = s.expandRows()
	*s = s.expandColumns()
}

func (g galaxyMap) expandRows() galaxyMap {
	nrOfRows := len(g)
	for i := 0; i < nrOfRows; i++ {
		hasGalaxy := false
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '#' {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			g = expandRow(g, i)
			i++
			nrOfRows++
		}
	}

	return g
}

func expandRow(g galaxyMap, index int) galaxyMap {
	expansion := make([]byte, len(g[index]))
	for j := 0; j < len(g[index]); j++ {
		expansion[j] = '.'
	}

	return insertRow(g, expansion, index)
}

func insertRow(g galaxyMap, row []byte, index int) galaxyMap {
	g = append(g, []byte{})
	copy(g[index+1:], g[index:])
	g[index] = row

	return g
}

func (g galaxyMap) expandColumns() galaxyMap {
	nrOfColumns := len(g[0])
	for i := 0; i < nrOfColumns; i++ {
		hasGalaxy := false
		for j := 0; j < len(g); j++ {
			if g[j][i] == '#' {
				hasGalaxy = true
				break
			}
		}

		if !hasGalaxy {
			g = expandColumn(g, i)
			i++
			nrOfColumns++
		}
	}

	return g
}

func expandColumn(g galaxyMap, index int) galaxyMap {
	for i, row := range g {
		g[i] = insertColumnInRow(row, '.', index)
	}

	return g
}

func insertColumnInRow(row []byte, char byte, index int) []byte {
	row = append(row, char)
	copy(row[index+1:], row[index:])
	row[index] = char

	return row
}

func parseInputToGalaxyMap(filePath string) galaxyMap {
	galaxyMap := galaxyMap{}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		galaxyMap = append(galaxyMap, []byte(scanner.Text()))
	}

	return galaxyMap
}
