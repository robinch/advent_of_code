package day15

type hashMap []*List

func newHashMap() hashMap {
	hashMap := make(hashMap, 256)
	for i := 0; i < 256; i++ {
		hashMap[i] = new(List)
	}

	return hashMap
}

func (h *hashMap) doStep(s step) {
	switch s.opertation {
	case '=':
		h.add(s.lense)
	case '-':
		h.remove(s.lense)
	}
}

func (h *hashMap) add(l lense) {
	(*h)[hash(l.label)].Add(l)
}

func (h *hashMap) remove(l lense) {
	(*h)[hash(l.label)].Remove(l)
}

func hash(s string) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash += int(s[i])
		hash *= 17
		hash %= 256
	}
	return hash
}
