package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
	input := parse_input("input.txt")
	left := input[0]
	right := input[1]

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += abs(left[i] - right[i])
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func Part2() {
	input := parse_input("input.txt")
	left := input[0]
	right := input[1]

	occurence := make(map[int]int)

	for _, value := range right {
		occurence[value]++
	}

	sum := 0

	for _, value := range left {
		sum += value * occurence[value]
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func parse_input(path string) [][]int {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for scanner.Scan() {
		seperator := "   "

		row := scanner.Text()
		numbers := strings.Split(row, seperator)

		left, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return [][]int{leftList, rightList}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	Part1()
	Part2()
}
