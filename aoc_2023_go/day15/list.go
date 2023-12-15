package day15

type List struct {
	head *node
	len  int
}

type node struct {
	next *node
	val  lense
}

func (l *List) Add(lense lense) {
	if l.head == nil {
		l.head = &node{val: lense}
		l.len++
		return
	}

	if l.head.val.label == lense.label {
		l.head.val.focalLength = lense.focalLength
		return
	}

	current := l.head
	for ; current.next != nil; current = current.next {
		if current.next.val.label == lense.label {
			current.next.val.focalLength = lense.focalLength
			return
		}
	}

	current.next = &node{val: lense}
	l.len++
}

func (l *List) Remove(lense lense) {
	if l.head == nil {
		return
	}

	current := l.head

	if current.val.label == lense.label {
		l.head = current.next
		l.len--
		return
	}

	for ; current.next != nil; current = current.next {
		if current.next.val.label == lense.label {
			current.next = current.next.next
			l.len--
			return
		}
	}
}

func (l *List) Len() int {
	return l.len
}
