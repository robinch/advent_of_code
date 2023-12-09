package day9

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("part1_test_input.txt")
	want := 114

	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("part2_test_input.txt")
	want := 5

	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}
