package day05

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("../test_inputs/day5.txt")
	want := 35

	if got != want {
		t.Errorf("Want %d but got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("../test_inputs/day5.txt")
	want := 46

	if got != want {
		t.Errorf("Want %d but got %d", want, got)
	}
}
