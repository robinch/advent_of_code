package day04

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("../test_inputs/day4.txt")
	expect := 13

	if got != expect {
		t.Errorf("got %d, expected %d", got, expect)
	}
}
