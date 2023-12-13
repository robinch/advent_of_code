package day12

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 21

	assertEqual(t, got, want)
}

func TestPart2(t *testing.T) {
	got := Part2("test_input.txt")
	want := 525152

	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
