package day22

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pos struct {
	x, y, z int
}

type brick struct {
	start, end   pos
	above, below []*brick
}

type posToBrickMap map[pos]*brick

func (b *brick) String() string {
	return fmt.Sprintf("%v-%v", b.start, b.end)
}

func Part1(filePath string) int {
	bricks := readInput(filePath)
	sort.Slice(bricks, func(i, j int) bool {
		return less(bricks[i].start, bricks[j].start)
	})

	posToBrick := posToBrick(bricks)
	fmt.Printf("Bricks: %v\n", bricks)
	fmt.Printf("Pos to brick: %v\n", posToBrick)

	return 0
}

func posToBrick(bricks []brick) posToBrickMap {
	m := map[pos]*brick{}

	for i := range bricks {
		b := &bricks[i]
		xStart, xEnd := order(b.start.x, b.end.x)
		yStart, yEnd := order(b.start.y, b.end.y)
		zStart, zEnd := order(b.start.z, b.end.z)

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				for z := zStart; z <= zEnd; z++ {
					m[pos{x, y, z}] = b
				}
			}
		}
	}

	return m
}

func above(bricks []brick, posToBrickMap posToBrickMap) {
	maxHeight := bricks[len(bricks)-1].end.z
	for _, b := range bricks {
		xStart, xEnd := order(b.start.x, b.end.x)
		yStart, yEnd := order(b.start.y, b.end.y)
		zStart, _ := order(b.start.z, b.end.z)

		zStart++

		found := false

		for z := zStart; z <= maxHeight; z++ {
			if found {
				break
			}

			for x := xStart; x <= xEnd; x++ {
				for y := yStart; y <= yEnd; y++ {
					if b, ok := posToBrickMap[pos{x, y, z}]; ok {
						if !contains(b.above, b) {
							b.above = append(b.above, b)
						}
						found = true
					}
				}
			}
		}
	}

}

func contains(bricks []*brick, b *brick) bool {
	for _, brick := range bricks {
		if brick == b {
			return true
		}
	}

	return false
}

func order(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func readInput(filePath string) []brick {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bricks := []brick{}

	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		row := scanner.Text()
		brickEnds := strings.Split(row, "~")
		// x,y,z
		endA := strings.Split(brickEnds[0], ",")
		endB := strings.Split(brickEnds[1], ",")
		posA := pos{toInt(endA[0]), toInt(endA[1]), toInt(endA[2])}
		posB := pos{toInt(endB[0]), toInt(endB[1]), toInt(endB[2])}

		var b brick

		if less(posA, posB) {
			b = brick{posA, posB, []*brick{}, []*brick{}}
		} else {
			b = brick{posB, posA, []*brick{}, []*brick{}}
		}

		bricks = append(bricks, b)
		n++
	}

	return bricks
}

func less(a, b pos) bool {
	if a.z < b.z {
		return true
	} else if a.z > b.z {
		return false
	} else if a.y < b.y {
		return true
	} else if a.y > b.y {
		return false
	} else if a.x < b.x {
		return true
	} else {
		return false
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
