package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	sum := 0

	file, err := os.Open("inputs/day1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sum += getNumber(scanner.Text())
	}

	return sum
}

func getNumber(line string) int {
	firstNumber := 0
	secondNumber := 0

	s := strings.Split(line, "")
	i := 0

	for i < len(line) {
		n, err := strconv.Atoi(s[i])

		if err == nil {
			firstNumber = n
			break
		}

		i++
	}

	j := len(line) - 1

	for j >= i {
		n, err := strconv.Atoi(s[j])

		if err == nil {
			secondNumber = n
			break
		}

		j--
	}

	return 10*firstNumber + secondNumber
}
