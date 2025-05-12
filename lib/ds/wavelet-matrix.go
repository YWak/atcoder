package ds

import (
	"math/bits"

	"github.com/ywak/atcoder/lib/math"
)

// https://judge.u-aizu.ac.jp/onlinejudge/review.jsp?rid=3370714#1

// WaveletMatrixは、整数の配列に対して以下の操作を効率的に行うデータ構造です。
type WaveletMatrix struct {
	bits     []SuccinctIndexableDictionary
	beginOne []int
	begin    map[int]int
	n        int
	maxElem  int
	bitsize  int
}

func NewWaveletMatrix(data []int) *WaveletMatrix {
	wm := WaveletMatrix{}
	wm.n = len(data)
	wm.begin = map[int]int{}
	for _, v := range data {
		if v > wm.maxElem {
			wm.maxElem = v
		}
	}
	wm.bitsize = math.Max(1, bits.Len64(uint64(wm.maxElem)))
	wm.beginOne = make([]int, wm.bitsize)
	for i := 0; i < wm.bitsize; i++ {
		wm.bits = append(wm.bits, *NewSuccinctIndexableDictionary(wm.n))
	}

	values, temp := make([]int, wm.n), make([]int, wm.n)
	copy(values, data)

	for i := 0; i < wm.bitsize; i++ {
		k := 0
		for j, v := range values {
			bit := (v >> (wm.bitsize - 1 - i)) & 1

			if bit == 0 {
				temp[k] = v
				k++
				wm.bits[i].Set(j, 0)
			}
		}
		wm.beginOne[i] = k

		for j, v := range values {
			bit := (v >> (wm.bitsize - 1 - i)) & 1
			if bit == 1 {
				temp[k] = v
				k++
				wm.bits[i].Set(j, 1)
			}
		}

		wm.bits[i].Build()
		values, temp = temp, values
	}
	for i := len(values) - 1; i >= 0; i-- {
		wm.begin[values[i]] = i
	}

	return &wm
}

// Getはwm[index]を返します。
func (wm *WaveletMatrix) Get(index int) int {
	if index < 0 || index >= wm.n {
		return -1
	}

	value := 0
	for i := 0; i < wm.bitsize; i++ {
		bit := wm.bits[i].Get(index)
		value = (value << 1) | bit

		index = wm.bits[i].Rank(bit, index)
		if bit == 1 {
			index += wm.beginOne[i]
		}
	}

	return value
}

// Selectはi番目のcの位置+1を返す。rankは1-origin。
func (wm *WaveletMatrix) Select(value, rank int) int {
	if rank <= 0 {
		return -1
	}

	index := wm.begin[value] + rank
	for i := 0; i < wm.bitsize; i++ {
		bit := (value >> i) & 1
		if bit == 1 {
			index -= wm.beginOne[wm.bitsize-i-1]
		}
		index = wm.bits[wm.bitsize-i-1].Select(bit, index)
	}

	return index
}

func (wm *WaveletMatrix) RangeFreq(left, right int, minVal, maxVal int) int {
	_, maxi, _ := wm.rankAll(maxVal, left, right)
	_, mini, _ := wm.rankAll(minVal, left, right)
	return maxi - mini
}

func (wm *WaveletMatrix) rankAll(value, left, right int) (int, int, int) {
	num := right - left
	if left >= right {
		return 0, 0, 0
	}
	if value > wm.maxElem || right == 0 {
		return 0, num, 0
	}
	less, more := 0, 0
	for i := 0; i < wm.bitsize && left < right; i++ {
		bit := (value >> (wm.bitsize - 1 - i)) & 1
		rank0left := wm.bits[i].Rank(0, left)
		rank0right := wm.bits[i].Rank(0, right)
		rank1left := left - rank0left
		rank1right := right - rank0right

		if bit == 1 {
			less += rank0right - rank0left
			left = wm.beginOne[i] + rank1left
			right = wm.beginOne[i] + rank1right
		} else {
			more += rank1right - rank1left
			left = rank0left
			right = rank0right
		}
	}

	rank := num - less - more
	return rank, less, more
}
