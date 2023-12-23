package day22

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetBrickPositions(t *testing.T) {
	brick := brick{pos{0, 0, 1}, pos{2, 0, 1}, nil, nil}
	got := getBrickPositions(brick)
	want := []pos{{0, 0, 1}, {1, 0, 1}, {2, 0, 1}}

	if !equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func equal(a, b []pos) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
