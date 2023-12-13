package day06

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("../test_inputs/day6.txt")
	want := 288

	if got != want {
		t.Errorf("Want %d but got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("../test_inputs/day6.txt")
	want := 71503

	if got != want {
		t.Errorf("Want %d but got %d", want, got)
	}
}

func TestGetHoldTime(t *testing.T) {
	got := getHoldTime(9, 7)
	want := 1.6972243622680054

	if got != want {
		t.Errorf("Want %g but got %g", want, got)
	}
}
