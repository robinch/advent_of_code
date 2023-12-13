package day04

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("../test_inputs/day4.txt")
	expect := 30

	if got != expect {
		t.Errorf("got %d, expected %d", got, expect)
	}
}
