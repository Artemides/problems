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

func (s *IntSet) Union(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
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
	set2.Add(62)
	set2.Add(149)
	set.Add(62)
	set.Add(63)
	set.Add(98)
	set.Add(182)
	set.Union(&set2)
	fmt.Println(set.String())
	fmt.Println(set.Len())
	set.Remove(62)
	fmt.Println(set.String())
	// set.Clear()
	set.Add(99)
	fmt.Println("Set: \t", set.String())
	set3 := set.Copy()
	set3.AddAll(133, 1564, 2312, 10)
	fmt.Println("Set3: \t", set3.String())

}
