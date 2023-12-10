package day10

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

type direction int

type pipe struct {
	pipeType  byte
	direction direction
	coord     coordinate
}

func (s pipe) String() string {
	return fmt.Sprintf("{%c, %q, %v}", s.pipeType, directionToString(s.direction), s.coord)
}

func directionToString(direction direction) string {
	switch direction {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	}

	return "UNKNOWN"
}

type coordinate struct {
	x int
	y int
}

type pipesMap [][]byte

func (p pipesMap) String() string {
	str := ""

	for _, line := range p {
		str += fmt.Sprintf("%s\n", string(line))
	}

	return str
}

func Part1(filePath string) int {
	pipesMap := getPipesMap(filePath)
	startingPoint, err := findStartingPoint(pipesMap)
	check(err)

	pipeLoop, err := pipeLoop(pipesMap, startingPoint)

	check(err)

	return len(pipeLoop) / 2
}

func Part2(filePath string) int {
	pipesMap := getPipesMap(filePath)
	startingPoint, err := findStartingPoint(pipesMap)
	check(err)

	pipeLoop, err := pipeLoop(pipesMap, startingPoint)
	check(err)

	updateSPipePipeType(pipeLoop)
	enclosedTiles := enclosedTiles(pipesMap, pipeLoop)

	for _, et := range enclosedTiles {
		pipesMap[et.y][et.x] = 'O'
	}

	return len(enclosedTiles)
}

func enclosedTiles(pipeMap pipesMap, pipeLoop []pipe) []coordinate {
	pipeTypeLookUpTable := pipeTypeLookUpTable(pipeLoop)
	enclosedTiles := []coordinate{}

	for y, pipeRow := range pipeMap {
		for x := range pipeRow {
			coord := coordinate{x: x, y: y}
			_, partOfLoop := pipeTypeLookUpTable[coord]
			if !partOfLoop {
				if enclosed := isEnclosed(coord, pipeTypeLookUpTable); enclosed {
					enclosedTiles = append(enclosedTiles, coord)
				}
			}
		}
	}

	return enclosedTiles
}

func isEnclosed(coord coordinate, lookUp map[coordinate]byte) bool {
	wallCollisions := 0
	recentCorner := ' '
	for i := coord.x; i >= 0; i-- {
		if lookUp[coordinate{x: i, y: coord.y}] == '|' {
			wallCollisions++
			recentCorner = ' '
		} else if lookUp[coordinate{x: i, y: coord.y}] == '7' {
			recentCorner = '7'
		} else if lookUp[coordinate{x: i, y: coord.y}] == 'L' {
			if recentCorner == '7' {
				wallCollisions++
				recentCorner = ' '
			}
		} else if lookUp[coordinate{x: i, y: coord.y}] == 'J' {
			recentCorner = 'J'
		} else if lookUp[coordinate{x: i, y: coord.y}] == 'F' {
			if recentCorner == 'J' {
				wallCollisions++
				recentCorner = ' '
			}
		}
	}

	return wallCollisions%2 == 1
}

func pipeTypeLookUpTable(pipeLooop []pipe) map[coordinate]byte {
	lookUpTable := map[coordinate]byte{}

	for _, pipe := range pipeLooop {
		lookUpTable[pipe.coord] = pipe.pipeType
	}

	return lookUpTable
}

func updateSPipePipeType(pipeLoop []pipe) {
	sPipe := &pipeLoop[len(pipeLoop)-1]
	beforeS := pipeLoop[len(pipeLoop)-2]
	afterS := pipeLoop[0]

	above, below := sortByY(beforeS, afterS)

	if above.coord.x == below.coord.x {
		sPipe.pipeType = '|'
	} else if above.coord.y == below.coord.y {
		sPipe.pipeType = '-'
	} else if above.coord.x == sPipe.coord.x &&
		below.coord.x < sPipe.coord.x {
		sPipe.pipeType = 'J'
	} else if above.coord.x == sPipe.coord.x &&
		sPipe.coord.x < below.coord.x {
		sPipe.pipeType = 'L'
	} else if below.coord.x == sPipe.coord.x &&
		above.coord.x < sPipe.coord.x {
		sPipe.pipeType = '7'
	} else if below.coord.x == sPipe.coord.x &&
		above.coord.x > sPipe.coord.x {
		sPipe.pipeType = 'F'
	}
}

func sortByY(pipe1, pipe2 pipe) (pipe, pipe) {
	if pipe1.coord.y < pipe2.coord.y {
		return pipe1, pipe2
	}

	return pipe2, pipe1
}

func pipeLoop(pipesMap pipesMap, startingPoint coordinate) ([]pipe, error) {
	pipes := []pipe{}

	currentPipe, err := findFirstPipe(pipesMap, startingPoint)
	check(err)
	pipes = append(pipes, currentPipe)
	ok := true

	for ok {
		currentPipe, ok = step(pipesMap, currentPipe)

		pipes = append(pipes, currentPipe)

		if currentPipe.pipeType == 'S' {
			return pipes, nil
		}
	}

	return pipes, errors.New("Could not find pipe loop")
}

func findFirstPipe(pipesMap pipesMap, startingPoint coordinate) (pipe, error) {
	for _, direction := range []direction{UP, DOWN, LEFT, RIGHT} {
		firstPipe, ok := step(
			pipesMap,
			pipe{direction: direction, pipeType: 'S', coord: startingPoint},
		)

		if ok {
			return firstPipe, nil
		}
	}

	return pipe{}, errors.New("Could not find First Pipee")
}

func step(pipesMap pipesMap, currentPipe pipe) (pipe, bool) {
	possbilePipeTypes := directionToPossiblePipeTypes(currentPipe.direction)
	nextX, nextY := nextXandY(currentPipe, currentPipe.direction)

	if nextY >= 0 && nextY < len(pipesMap) &&
		nextX >= 0 && nextX < len(pipesMap[0]) {
		for _, pipeType := range possbilePipeTypes {
			if pipesMap[nextY][nextX] == pipeType {
				return pipe{
						pipeType:  pipeType,
						direction: nextDirection(currentPipe.direction, pipeType),
						coord:     coordinate{x: nextX, y: nextY},
					},
					true
			}
		}
	}

	return pipe{}, false
}

func nextXandY(currentPipe pipe, direction direction) (int, int) {
	switch direction {
	case UP:
		return currentPipe.coord.x, currentPipe.coord.y - 1
	case DOWN:
		return currentPipe.coord.x, currentPipe.coord.y + 1
	case LEFT:
		return currentPipe.coord.x - 1, currentPipe.coord.y
	default:
		return currentPipe.coord.x + 1, currentPipe.coord.y
	}
}

func nextDirection(currentDirection direction, pipeType byte) direction {
	switch currentDirection {
	case UP:
		switch pipeType {
		case '|':
			return UP
		case '7':
			return LEFT
		case 'F':
			return RIGHT
		}
	case DOWN:
		switch pipeType {
		case '|':
			return DOWN
		case 'L':
			return RIGHT
		case 'J':
			return LEFT
		}
	case LEFT:
		switch pipeType {
		case '-':
			return LEFT
		case 'L':
			return UP
		case 'F':
			return DOWN
		}
	case RIGHT:
		switch pipeType {
		case '-':
			return RIGHT
		case '7':
			return DOWN
		case 'J':
			return UP
		}
	}

	return -1
}

func directionToPossiblePipeTypes(direction direction) []byte {
	switch direction {
	case UP:
		return []byte{'|', '7', 'F', 'S'}
	case DOWN:
		return []byte{'|', 'L', 'J', 'S'}
	case LEFT:
		return []byte{'-', 'L', 'F', 'S'}
	case RIGHT:
		return []byte{'-', '7', 'J', 'S'}
	}

	return []byte{}
}

func getPipesMap(filePath string) pipesMap {
	pipesMap := pipesMap{}
	file, err := os.Open(filePath)

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		row := []byte(scanner.Text())
		pipesMap = append(pipesMap, row)
		i++
	}

	return pipesMap
}

func findStartingPoint(pipesMap pipesMap) (coordinate, error) {
	for i, pipeRow := range pipesMap {
		for j, currentPipe := range pipeRow {
			if currentPipe == 'S' {
				return coordinate{x: j, y: i}, nil
			}
		}
	}
	return coordinate{}, errors.New("could not find starting pipe")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
