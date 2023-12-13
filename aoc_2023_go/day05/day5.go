package day05

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mapping struct {
	destStart   int
	sourceStart int
	rangeLength int
}

func Part1(filePath string) int {
	minLocation := math.MaxInt32
	seeds, mappings := parseInput(filePath)

	for _, seed := range seeds {
		location := seedToLocation(seed, mappings)
		if location < minLocation {
			minLocation = location
		}
	}

	// debugPrint(seeds, mappings, locations)

	return minLocation
}

func Part2(filePath string) int {
	minLocation := math.MaxInt32
	seeds, mappings := parseInput(filePath)

	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < (seeds[i] + seeds[i+1]); seed++ {
			location := seedToLocation(seed, mappings)
			if location < minLocation {
				minLocation = location
			}
		}
	}

	// debugPrint(seeds, mappings, locations)

	return minLocation
}

func seedToLocation(seed int, mappings [][]mapping) int {
	source := seed
	for i := 0; i < len(mappings); i++ {
		source = getDestination(source, mappings[i])
	}

	return source
}

func getDestination(source int, mappings []mapping) int {
	for _, m := range mappings {
		if m.sourceStart <= source && source < m.sourceStart+m.rangeLength {
			diff := source - m.sourceStart

			return m.destStart + diff
		}
	}

	return source
}

func min(a []int) int {
	min := math.MaxInt32
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}

	return min
}

func parseInput(filePath string) ([]int, [][]mapping) {
	seeds := []int{}
	mappings := [][]mapping{}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seeds = getSeeds(line)
		}

		if strings.HasSuffix(line, "map:") {
			mappings = append(mappings, getMapping(scanner))
		}
	}

	return seeds, mappings
}

func getSeeds(line string) []int {
	split := strings.Split(line, ":")
	return getNumbers(split[1])
}

func getMapping(scanner *bufio.Scanner) []mapping {
	mappings := []mapping{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		numbers := getNumbers(line)
		mappings = append(mappings, mapping{
			destStart:   numbers[0],
			sourceStart: numbers[1],
			rangeLength: numbers[2]})
	}

	return mappings
}

func getNumbers(s string) []int {
	numbers := []int{}
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	matches := re.FindAllString(s, -1)
	for _, match := range matches {
		n, err := strconv.Atoi(match)

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}
	return numbers
}

func debugPrint(seeds []int, mappings [][]mapping, locations []int) {
	fmt.Printf("Seeds: %v\n", seeds)
	fmt.Printf("seed-to-soil mapping: %v\n", mappings[0])
	fmt.Printf("soil-to-fertilizer mapping: %v\n", mappings[1])
	fmt.Printf("fertilizer-to-water mapping: %v\n", mappings[2])
	fmt.Printf("water-to-light mapping: %v\n", mappings[3])
	fmt.Printf("light-to-temperature mapping: %v\n", mappings[4])
	fmt.Printf("temperature-to-humidity mapping: %v\n", mappings[5])
	fmt.Printf("humidity-to-location mapping: %v\n", mappings[6])
	fmt.Printf("locations %v\n", locations)
}
