package day10

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Run("test_input_2.txt", func(t *testing.T) {
		got := Part2("test_input_2.txt")
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test_input_3.txt", func(t *testing.T) {

		got := Part2("test_input_3.txt")
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test_input_4.txt", func(t *testing.T) {

		got := Part2("test_input_4.txt")
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
