package day22

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
