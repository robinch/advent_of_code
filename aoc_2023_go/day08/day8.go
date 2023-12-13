package day08

import (
	"bufio"
	"os"
	"regexp"
)

type Node struct {
	left  string
	right string
}

func Part1(filePath string) int {
	instructions, network := getInputs(filePath)
	goalFunction := func(name string) bool { return name == "ZZZ" }

	return stepsToGoal("AAA", goalFunction, instructions, network)
}

func Part2(filePath string) int {
	// Seems like all paths loops after finding the first goal
	// Can use LCD on all steps for each path
	instructions, network := getInputs(filePath)
	names := startingNames(network)
	goalFunction := func(name string) bool { return endsWith(name, 'Z') }

	steps := []int{}
	for _, name := range names {
		steps = append(steps, stepsToGoal(name, goalFunction, instructions, network))
	}

	return lcd(steps)
}

func stepsToGoal(start string, goalFunction func(name string) bool, instructions []byte, network map[string]Node) int {
	i := 0
	steps := 0
	name := start
	nrOfInstructions := len(instructions)

	for true {
		if goalFunction(name) {
			break
		}

		if instructions[i] == 'L' {
			name = network[name].left
		} else {
			name = network[name].right
		}

		steps++
		i = (i + 1) % nrOfInstructions
	}

	return steps
}

func startingNames(network map[string]Node) []string {
	startingNames := []string{}
	for key := range network {
		if endsWith(key, 'A') {
			startingNames = append(startingNames, key)
		}
	}

	return startingNames
}

func endsWith(name string, r byte) bool {
	if name[2] == r {
		return true
	}

	return false
}

func getInputs(filePath string) ([]byte, map[string]Node) {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := []byte(scanner.Text())

	network := map[string]Node{}

	scanner.Scan()

	for scanner.Scan() {
		name, node := getNameAndNode(scanner.Text())
		network[name] = node
	}

	return instructions, network
}

func getNameAndNode(line string) (string, Node) {
	re := regexp.MustCompile(`[A-Z0-9]+`)
	matches := re.FindAllString(line, -1)
	name := matches[0]
	node := Node{left: matches[1], right: matches[2]}

	return name, node
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcd(nums []int) int {
	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}
