package day11

import "testing"


func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 374

	if got != want {
		t.Errorf("Part1() = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("test_input.txt", 100)
	want := 8410

	if got != want {
		t.Errorf("Part1() = %d; want %d", got, want)
	}
}
