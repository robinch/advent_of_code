package day21

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt", 6)
	want := 16

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestQueue(t *testing.T) {
	q := queue{}

	_, ok := q.dequeue()

	if ok != false {
		t.Error("Expected ok to be false")
	}

	n1 := node{0, 0, 0}
	n2 := node{1, 2, 0}

	q.enqueue(n1)
	q.enqueue(n2)

	p, ok := q.dequeue()

	if ok != true {
		t.Error("Expected ok to be true")
	}

	if p != n1 {
		t.Errorf("Expected p to be {1, 2}, got %v", p)
	}

	p, ok = q.dequeue()
	
	if ok != true {
		t.Error("Expected ok to be true")
	}

	if p != n2 {
		t.Errorf("Expected p to be {0, 0}, got %v", p)
	}
}
