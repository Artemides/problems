package methods

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (set *IntSet) AddAll(ints ...int) {
	for _, num := range ints {
		set.Add(num)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			continue
		}
		s.words[i] &= tword

	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			continue
		}
		s.words[i] &^= tword

	}
}

func (s *IntSet) SymetricDiferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, tword)
			continue
		}
		s.words[i] ^= tword

	}
}

func (s *IntSet) Union(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
			continue
		}

		s.words = append(s.words, tword)

	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if !s.Has(x) {
		return
	}

	s.words[word] ^= 1 << bit

}

func (set *IntSet) Len() uint {
	var len uint
	for _, word := range set.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<j) == 0 {
				continue
			}
			len++
		}
	}
	return len
}

func (set *IntSet) Elems() []int {
	len := set.Len()
	elems := make([]int, len)
	elIdx := 0
	for i, word := range set.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<j) == 0 {
				continue
			}
			elems[elIdx] = i*64 + j
			elIdx++
		}
	}

	return elems
}

func (set *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range set.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<j) == 0 {
				continue
			}

			if buf.Len() > 1 {
				buf.WriteByte(' ')
			}
			buf.WriteString(fmt.Sprintf("%d", i*64+j))
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (set *IntSet) Clear() {
	set.words = set.words[:0]
}
func (set *IntSet) Copy() *IntSet {
	newSet := IntSet{
		words: make([]uint64, len(set.words), cap(set.words)),
	}
	copy(newSet.words, set.words)

	return &newSet
}

func RunSet() {
	var set IntSet

	set.AddAll(65, 11, 76, 112, 165)
	fmt.Println("Set 1: ", set.String())

	setElems := set.Elems()
	fmt.Println("Set 1 Elems: ", setElems)
}
