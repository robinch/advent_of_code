package main

import (
	"aoc_2023_go/day1"
	"aoc_2023_go/day2"
	"aoc_2023_go/day3"
	"aoc_2023_go/day4"
	"fmt"
)

func main() {
	fmt.Printf("Day 1, part 1: %d\n", day1.Part1())
	fmt.Printf("Day 1, part 2: %d\n", day1.Part2())

	fmt.Printf("Day 2, part 1: %d\n", day2.Part1())
	fmt.Printf("Day 2, part 2: %d\n", day2.Part2())

	fmt.Printf("Day 3, part 1: %d\n", day3.Part1())
	fmt.Printf("Day 3, part 2: %d\n", day3.Part2())
  
	fmt.Printf("Day 4, part 1: %d\n", day4.Part1("inputs/day4.txt"))
}
