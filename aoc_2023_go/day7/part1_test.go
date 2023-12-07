package day7

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("../test_inputs/day7.txt")
	want := 6440

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("../test_inputs/day7.txt")
	want := 5905

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGetHandVal(t *testing.T) {
	got := getHandVal([]rune{'A', '3', '4', '5', '2'}, false)
	want := 719852

	if got != want {
		t.Errorf("got %d, want %d\n", got, want)
	}
}

func TestGetType(t *testing.T) {
	got := getType([]rune{'A', 'A', 'Q', 'Q', 'A'}, false)
	want := fullHouse

	if got != want {
		t.Errorf("got %d, want %d\n", got, want)
	}
}
