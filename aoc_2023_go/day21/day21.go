package day21

import (
	"bufio"
	"os"
)

type puzzle [][]byte

func (p puzzle) String() string {
	s := ""
	for _, row := range p {
		s += string(row) + "\n"
	}

	return s
}

type node struct {
	row, col, depth int
}

type queue []node

type pos struct {
	row, col int
}

func (q *queue) enqueue(p node) {
	*q = append(*q, p)
}

func (q *queue) dequeue() (node, bool) {
	if len(*q) == 0 {
		return node{}, false
	}

	p := (*q)[0]
	*q = (*q)[1:]
	return p, true
}

func Part1(filePath string, steps int) int {
	puzzle := readInput(filePath)

	sPos := getStartPos(puzzle)

	counter := 0

	q := queue{}
	visited := map[pos]struct{}{}
	q.enqueue(sPos)

	for len(q) > 0 {
		n, _ := q.dequeue()

		if _, ok := visited[pos{n.row, n.col}]; ok {
			continue
		}

		visited[pos{n.row, n.col}] = struct{}{}
		
		if n.depth %2 == 0 {
			counter++
		}
		
		if n.depth == steps {
			continue
		}

		// above
		if n.row - 1 >= 0 && puzzle[n.row - 1][n.col] != '#' {
			q.enqueue(node{row: n.row - 1, col: n.col, depth: n.depth + 1})
		}

		// left
		if n.col - 1 >= 0 && puzzle[n.row][n.col - 1] != '#' {
			q.enqueue(node{row: n.row, col: n.col - 1, depth: n.depth + 1})
		}

		// below
		if n.row + 1 < len(puzzle) && puzzle[n.row + 1][n.col] != '#' {
			q.enqueue(node{row: n.row + 1, col: n.col, depth: n.depth + 1})
		}

		// right
		if n.col + 1 < len(puzzle[0]) && puzzle[n.row][n.col + 1] != '#' {
			q.enqueue(node{row: n.row, col: n.col + 1, depth: n.depth + 1})
		}
	}

	return counter
}

func copyPuzzle(p puzzle) puzzle {
	copiedPuzzle := make(puzzle, len(p))

	for i := 0; i < len(p); i++ {
		copiedPuzzle[i] = make([]byte, len(p[i]))
		copy(copiedPuzzle[i], p[i])
	}

	return copiedPuzzle
}

func readInput(filePath string) puzzle {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	puzzle := puzzle{}

	for scanner.Scan() {
		puzzle = append(puzzle, []byte(scanner.Text()))
	}

	return puzzle
}

func getStartPos(p puzzle) node {
	for row := 0; row < len(p); row++ {
		for col := 0; col < len(p[0]); col++ {
			if p[row][col] == 'S' {
				return node{row: row, col: col, depth: 0}
			}
		}
	}
	panic("No start pos found")
}
