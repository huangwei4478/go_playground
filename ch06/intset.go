package ch06

type Inset struct {
	words []uint64
}

func (s *Inset) Has(x int) bool {
	word, bit := x / 64, uint(x % 64)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *Inset) Add(x int) {
	word, bit := x / 64, uint(x % 64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}


