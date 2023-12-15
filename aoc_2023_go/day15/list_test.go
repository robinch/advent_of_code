package day15

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add to empty List", func(t *testing.T) {
		l := List{}
		l.Add(lense{label: "test", focalLength: 1})
		got := l.head.val

		want := lense{label: "test", focalLength: 1}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Add to non-empty List", func(t *testing.T) {
		l := List{}
		l.Add(lense{label: "test", focalLength: 1})
		l.Add(lense{label: "test2", focalLength: 2})
		got := l.head.next.val

		want := lense{label: "test2", focalLength: 2}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Add to existing lense List", func(t *testing.T) {
		l := List{}
		l.Add(lense{label: "test", focalLength: 1})
		l.Add(lense{label: "test2", focalLength: 3})
		l.Add(lense{label: "test", focalLength: 2})
		got := l.head.val

		want := lense{label: "test", focalLength: 2}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("Remove from empty List", func(t *testing.T) {
		l := List{}
		l.Remove(lense{label: "test", focalLength: 1})
		got := l.head

		if got != nil {
			t.Errorf("got %v want %v", got, nil)
		}
	})
	t.Run("Remove from non-empty List", func(t *testing.T) {
		l := List{}
		l.Add(lense{label: "test", focalLength: 1})
		l.Add(lense{label: "test2", focalLength: 2})
		l.Remove(lense{label: "test", focalLength: 1})
		got := l.head.val

		want := lense{label: "test2", focalLength: 2}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Remove last element from non-empty List", func(t *testing.T) {
		l := List{}
		l.Add(lense{label: "test", focalLength: 1})
		l.Add(lense{label: "test2", focalLength: 2})
		l.Remove(lense{label: "test2", focalLength: 2})
		got := l.head.val

		want := lense{label: "test", focalLength: 1}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestLen(t *testing.T) {
	t.Run("Empty List", func(t *testing.T) {
		l := List{}
		got := l.Len()

		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Non-empty List", func(t *testing.T) {
		l := List{}
		l.Add(lense{label: "test", focalLength: 1})
		l.Add(lense{label: "test2", focalLength: 2})
		got := l.Len()

		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
