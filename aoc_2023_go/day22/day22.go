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
	label        string
	above, below []*brick
}

type posToBrickMap map[pos]*brick

func (b *brick) String() string {
	return fmt.Sprintf("%v%s%v", b.start, b.label, b.end)
}

func (b *brick) addBelow(brick *brick) {
	b.below = append(b.below, brick)
}

func (b *brick) containsBelow(brick *brick) bool{
	return contains(b.below, brick)
}

func (b *brick) addAbove(brick *brick) {
	b.above = append(b.above, brick)
}

func (b *brick) containsAbove(brick *brick) bool{
	return contains(b.above, brick)
}


func Part1(filePath string) int {
	bricks := readInput(filePath)
	sort.Slice(bricks, func(i, j int) bool {
		return less(bricks[i].start, bricks[j].start)
	})

	posToBrick := createPosToBrickMap(bricks)
	pushDown(bricks, posToBrick)
	populateAbove(bricks)

	bricksToDisintegrate := []brick{}

	for _, b := range bricks {
		disintegrate := true
		for _, above := range b.above {
			if len(above.below) == 1 {
				disintegrate = false
				break
			}
		}

		if disintegrate {
			bricksToDisintegrate = append(bricksToDisintegrate, b)
		}
	}

	return len(bricksToDisintegrate)
}

func createPosToBrickMap(bricks []brick) posToBrickMap {
	m := map[pos]*brick{}

	for i := range bricks {
		b := &bricks[i]
		for _, p := range getBrickPositions(*b) {
			m[p] = b
		}
	}

	return m
}

func getBrickPositions(b brick) []pos {
	positions := []pos{}

	xStart, xEnd := order(b.start.x, b.end.x)
	yStart, yEnd := order(b.start.y, b.end.y)
	zStart, zEnd := order(b.start.z, b.end.z)

	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			for z := zStart; z <= zEnd; z++ {
				positions = append(positions, pos{x, y, z})
			}
		}
	}

	return positions
}

func pushDown(bricks []brick, posToBrickMap posToBrickMap) {
	for i := range bricks {
		b := &bricks[i]

		zStart, _ := order(b.start.z, b.end.z)

		collision := false

		positions := getBrickPositions(*b)

		for z := zStart; z > 0; z-- {
			if collision {
				break
			}

			for _, p := range positions {
				if below, ok := posToBrickMap[pos{p.x, p.y, z-1}]; ok {
					if !b.containsBelow(below) {
						b.addBelow(below)
					}
					updateBrickZ(b, posToBrickMap, z)

					collision = true
				}
			}
		}
	}
}

func updateBrickZ(b *brick, posToBrickMap posToBrickMap, newZ int) {
		positions := getBrickPositions(*b)
		for _, p := range positions {
			delete(posToBrickMap, p)
			posToBrickMap[pos{p.x, p.y, newZ}] = b
		}
		
		b.start.z = newZ
		b.end.z = newZ
}


func populateAbove(bricks []brick) {
	for i := range bricks {
		b := &bricks[i]

		for i := range b.below {
			below := b.below[i]
			below.addAbove(b)
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

		label := fmt.Sprintf("%c", 'A'+n)

		if less(posA, posB) {
			b = brick{posA, posB, label, []*brick{}, []*brick{}}
		} else {
			b = brick{posB, posA, label, []*brick{}, []*brick{}}
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
