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

func Part1(filePath string) int {
	bricks := readInput(filePath)
	sort.Slice(bricks, func(i, j int) bool {
		return less(bricks[i].start, bricks[j].start)
	})

	posToBrick := createPosToBrickMap(bricks)

	fmt.Println("Bricks:")
	for _, b := range bricks {
		fmt.Printf("%s %v-%v\n", b.label, b.start, b.end)
	}


	// fmt.Printf("Pos toBricks\n brick: %v\n", posToBrick)

	// above(bricks, posToBrick)

	pushDown(bricks, posToBrick)
	fmt.Println("After push down:")
	for _, b := range bricks {
		fmt.Printf("%s %v-%v: %v\n", b.label, b.start, b.end, b.above)
	}

	return 0
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

		stop := false

		positions := getBrickPositions(*b)

		for z := zStart; z > 0; z-- {
			if stop {
				break
			}

			for _, p := range positions {
				if below, ok := posToBrickMap[pos{p.x, p.y, z-1}]; ok {
					fmt.Printf("%s found %s below\n", b.label, below.label)
					if !contains(b.below, below) {
						b.below = append(b.below, below)
					}
					updateBrickZ(b, posToBrickMap, z)

					stop = true
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

// func above(bricks []brick, posToBrickMap posToBrickMap) {
// 	maxHeight := bricks[len(bricks)-1].end.z
// 	for _, b := range bricks {
// 		xStart, xEnd := order(b.start.x, b.end.x)
// 		yStart, yEnd := order(b.start.y, b.end.y)
// 		zStart, _ := order(b.start.z, b.end.z)
//
// 		zStart++
//
// 		found := false
//
// 		for z := zStart; z <= maxHeight; z++ {
// 			if found {
// 				break
// 			}
//
// 			for x := xStart; x <= xEnd; x++ {
// 				for y := yStart; y <= yEnd; y++ {
// 					if above, ok := posToBrickMap[pos{x, y, z}]; ok {
// 						fmt.Printf("%s above %s\n", above.label, b.label)
// 						if !contains(b.above, above) {
// 							b.above = append(b.above, above)
// 						}
// 						found = true
// 					}
// 				}
// 			}
// 		}
// 	}
//
// }

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
