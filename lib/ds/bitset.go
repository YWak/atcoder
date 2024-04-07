package ds

import "math/bits"

type BitSet struct {
	size int
	bits []uint64
}

// NewBitSetはsizeだけのboolを管理できるBitSetを作ります。
func NewBitSet(size int) *BitSet {
	return &BitSet{
		size: size,
		bits: make([]uint64, (size+64-1)/64),
	}
}

// Setはindex番目のboolをvalueにします。
func (bs *BitSet) Set(index int, value bool) {
	if value {
		bs.bits[index/64] |= 1 << (index % 64)
	} else {
		bs.bits[index/64] ^= 1 << (index % 64)
	}
}

// Getはindex番目のboolを取得します。
func (bs *BitSet) Get(index int) bool {
	return ((bs.bits[index/64] >> (index % 64)) & 1) == 1
}

// Countはtrueになっている要素を数えて返します。
func (bs *BitSet) Count() int {
	ans := 0
	for _, v := range bs.bits {
		ans += bits.OnesCount64(v)
	}
	return ans
}

// Clearは全boolをfalseにします。
func (bs *BitSet) Clear() {
	for i := range bs.bits {
		bs.bits[i] = 0
	}
}

// Andは渡されたBitSetとの論理積を設定します。
func (bs *BitSet) And(other *BitSet) {
	if bs.size != other.size {
		panic("size not matched")
	}

	for i := range bs.bits {
		bs.bits[i] &= other.bits[i]
	}
}

// Orは渡されたBitSetとの論理和を設定します。
func (bs *BitSet) Or(other *BitSet) {
	if bs.size != other.size {
		panic("size not matched")
	}

	for i := range bs.bits {
		bs.bits[i] |= other.bits[i]
	}
}

// Xorは渡されたBitSetとの排他的論理和を設定します。
func (bs *BitSet) Xor(other *BitSet) {
	if bs.size != other.size {
		panic("size not matched")
	}
	for i := range bs.bits {
		bs.bits[i] ^= other.bits[i]
	}
}
