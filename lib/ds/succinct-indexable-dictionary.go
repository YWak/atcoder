package ds

import "math/bits"

// 完備辞書
type SuccinctIndexableDictionary struct {
	length int
	bits   []uint32
	sum    []uint32
}

func NewSuccinctIndexableDictionary(length int) *SuccinctIndexableDictionary {
	blocks := (length + 31) >> 5
	return &SuccinctIndexableDictionary{
		length: length,
		bits:   make([]uint32, blocks),
		sum:    make([]uint32, blocks),
	}
}

func (s *SuccinctIndexableDictionary) Set(i int) {
	s.bits[i>>5] |= 1 << uint(i&31)
}

func (s *SuccinctIndexableDictionary) Build() {
	s.sum[0] = 0
	for i := 1; i < len(s.sum); i++ {
		s.sum[i] = s.sum[i-1] + uint32(bits.OnesCount32(s.bits[i-1]))
	}
}

func (s *SuccinctIndexableDictionary) Get(i int) int {
	return int(s.bits[i>>5]>>uint(i&31)) & 1
}

func (s *SuccinctIndexableDictionary) Rank(i int) int {
	return (int(s.sum[i>>5]) + bits.OnesCount32(s.bits[i>>5]&((1<<uint(i&31))-1)))
}

func (s *SuccinctIndexableDictionary) Rank0(i int) int {
	return i - s.Rank(i)
}

func (s *SuccinctIndexableDictionary) Rank1(i int) int {
	return s.Rank(i)
}
