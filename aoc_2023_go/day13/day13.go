package day13

import (
	"bufio"
	"fmt"
	"os"
)

type terrain [][]byte

func (t terrain) String() string {
	str := ""

	for _, line := range t {
		str += fmt.Sprintf("%s\n", string(line))
	}

	str += "\n"

	return str
}

func Part1(filePath string) int {
	sum := 0
	terrains := pareInput(filePath)
	for _, terrain := range terrains {
		col, vFound := verticalMirror(terrain, false)
		row, hFound := horizontalMirror(terrain, false)

		if vFound {
			sum += col
		}
		if hFound {
			sum += row * 100
		}
	}
	return sum
}

func Part2(filePath string) int {
	sum := 0
	terrains := pareInput(filePath)
	for _, terrain := range terrains {
		col, vFound := verticalMirror(terrain, true)
		row, hFound := horizontalMirror(terrain, true)

		if vFound {
			sum += col
		}
		if hFound {
			sum += row * 100
		}
	}
	return sum
}

func verticalMirror(terrain terrain, deSmudge bool) (int, bool) {
	for col := 0; col < len(terrain[0])-1; col++ {
		mirrorFound, hasDesmuged := checkVertically(col, col+1, terrain, deSmudge)

		if deSmudge && mirrorFound && hasDesmuged || !deSmudge && mirrorFound {
			return col + 1, true
		}
	}

	return 0, false
}

func checkVertically(startLeft, startRight int, terrain terrain, deSmudge bool) (bool, bool) {
	hasDesmuged := false
	for row := 0; row < len(terrain); row++ {
		left := startLeft
		right := startRight

		for left >= 0 && right < len(terrain[0]) {
			if terrain[row][left] != terrain[row][right] {
				if !deSmudge || deSmudge && hasDesmuged {
					return false, false
				}

				hasDesmuged = true
			}

			left--
			right++
		}
	}

	return true, hasDesmuged
}

func horizontalMirror(terrain terrain, desmudge bool) (int, bool) {
	for row := 0; row < len(terrain)-1; row++ {
		mirrorFound, hasDesmuged := checkHorizontally(row, row+1, terrain, desmudge)

		if desmudge && mirrorFound && hasDesmuged || !desmudge && mirrorFound {
			return row + 1, true
		}
	}

	return 0, false
}

func checkHorizontally(startTop, startBottom int, terrain terrain, deSmudge bool) (bool, bool) {
	hasDesmuged := false
	for col := 0; col < len(terrain[0]); col++ {
		top := startTop
		bottom := startBottom

		for top >= 0 && bottom < len(terrain) {
			if terrain[top][col] != terrain[bottom][col] {
				if !deSmudge || deSmudge && hasDesmuged {
					return false, false
				}

				hasDesmuged = true
			}

			top--
			bottom++
		}
	}

	return true, hasDesmuged
}

func pareInput(filePath string) []terrain {
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	terrains := []terrain{}
	input := terrain{}

	for scanner.Scan() {
		row := []byte(scanner.Text())

		if len(row) == 0 {
			terrains = append(terrains, input)
			input = terrain{}
		} else {
			input = append(input, row)
		}
	}

	terrains = append(terrains, input)

	return terrains
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
