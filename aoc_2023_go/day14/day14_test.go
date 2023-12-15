package day14

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 136

	if got != want {
		t.Errorf("Part1() = %d; want %d", got, want)
	}
}


func TestPart2(t *testing.T) {
	got := Part2("test_input.txt")
	want := 64

	if got != want {
		t.Errorf("Part2() = %d; want %d", got, want)
	}
}

func TestTiltNorth(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltNorth(puzzleInput)

	want := `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
`


	checkTilt(t, puzzleInput, want)
}

func TestTiltSouth(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltSouth(puzzleInput)

	want :=`.....#....
....#....#
...O.##...
...#......
O.O....O#O
O.#..O.#.#
O....#....
OO....OO..
#OO..###..
#OO.O#...O
`


	checkTilt(t, puzzleInput, want)

}

func TestTiltWest(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltWest(puzzleInput)

	want :=`O....#....
OOO.#....#
.....##...
OO.#OO....
OO......#.
O.#O...#.#
O....#OO..
O.........
#....###..
#OO..#....
`

	checkTilt(t, puzzleInput, want)
}

func TestTiltEast(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltEast(puzzleInput)

	want := `....O#....
.OOO#....#
.....##...
.OO#....OO
......OO#.
.O#...O#.#
....O#..OO
.........O
#....###..
#..OO#....
`

	checkTilt(t, puzzleInput, want)
}

func TestOneCycle(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltAllDirections(puzzleInput)
	
	want := `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....
`

	checkTilt(t, puzzleInput, want)
}

func TestTwoCycle(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltAllDirections(puzzleInput)
	tiltAllDirections(puzzleInput)
	
	want := `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O
`

	checkTilt(t, puzzleInput, want)
}


func TestThreeCycle(t *testing.T) {
	puzzleInput := parseInput("test_input.txt")
	tiltAllDirections(puzzleInput)
	tiltAllDirections(puzzleInput)
	tiltAllDirections(puzzleInput)
	
	want := `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O
`

	checkTilt(t, puzzleInput, want)
}

func checkTilt(t *testing.T, puzzleInput puzzleInput, want string) {
	t.Helper()
	got := puzzleInput.String()

	if got != want {
		t.Errorf("\ngot\n%s\nwant\n%s", got, want)
	}
}
