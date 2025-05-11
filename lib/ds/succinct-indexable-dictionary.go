package ds

import (
	"math/bits"

	"github.com/ywak/atcoder/lib/math"
)

const LEVEL_L = 65536
const LEVEL_S = 256

// 完備辞書
type SuccinctIndexableDictionary struct {
	size  int
	bits  []uint64
	large []uint64
	small []uint16
	ones  int
}

func NewSuccinctIndexableDictionary(size int) *SuccinctIndexableDictionary {
	s := (size+64-1)/64 + 1

	return &SuccinctIndexableDictionary{
		size:  size,
		bits:  make([]uint64, s),
		large: make([]uint64, s/LEVEL_L+1),
		small: make([]uint16, s/LEVEL_S+1),
	}
}

// Setはi番目のビットをbitにします。
func (s *SuccinctIndexableDictionary) Set(i int, bit int) {
	pos, offset := s.index(i)

	if bit == 1 {
		s.bits[pos] |= 1 << offset
	} else {
		s.bits[pos] &= (^(1 << offset))
	}
}

func (s *SuccinctIndexableDictionary) Get(i int) int {
	pos, offset := s.index(i)
	return int(s.bits[pos]>>offset) & 1
}

func (s *SuccinctIndexableDictionary) index(i int) (int, int) {
	return i / 64, i % 64
}

func (s *SuccinctIndexableDictionary) Build() {
	num := 0
	for j, b := range s.bits {
		i := j * 64
		if i%LEVEL_L == 0 {
			s.large[i/LEVEL_L] = uint64(num)
		}
		if i%LEVEL_S == 0 {
			s.small[i/LEVEL_S] = uint16(num - int(s.large[i/LEVEL_L]))
		}
		num += bits.OnesCount64(b)
	}
	s.ones = num
}

// Rankは[0, i)に含まれるbitの数を返す。
func (s *SuccinctIndexableDictionary) Rank(bit, i int) int {
	if bit == 1 {
		return s.rank1(i)
	} else {
		return i - s.rank1(i)
	}
}

// [0, i)の1の数を返す。
func (s *SuccinctIndexableDictionary) rank1(i int) int {
	// LEVEL_L単位 + LEVEL_S単位 + 64bit単位 + 残り

	rank := s.large[i/LEVEL_L] + uint64(s.small[i/LEVEL_S])
	begin := (i / LEVEL_S) * LEVEL_S / 64
	end := i / 64

	for t := begin; t < end; t++ {
		rank += uint64(bits.OnesCount64(s.bits[t]))
	}
	rem := (i % LEVEL_S) % 64
	rank += uint64(bits.OnesCount64(s.bits[end] & ((1 << rem) - 1)))

	return int(rank)
}

// rank番目のbitの位置 + 1を返す。rankは1-origin。
func (s *SuccinctIndexableDictionary) Select(bit, rank int) int {
	// 暫定対応でrankを使って二分探索する
	// s.Rank(bit, ok+1) == rankとなるような最小のokを求める
	ng, ok := 0, s.size+1
	for math.Abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if s.Rank(bit, mid) >= rank {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}
