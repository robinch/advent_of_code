package day14

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 136

	if got != want {
		t.Errorf("Part1() = %d; want %d", got, want)
	}
}

