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
	var set2 IntSet
	set.AddAll(65, 11, 76, 112, 165)
	fmt.Println("Set 1: ", set.String())
	set2.AddAll(45, 76, 11, 435, 236)
	fmt.Println("Set 2: ", set2.String())
	set.IntersectWith(&set2)
	fmt.Println("S1 & S2:  ", set.String())
}
