package algorithm

import "math/rand"

type RollingHash struct {
	InitInts   func(arr []int)
	InitString func(value string)
	InitBytes  func(arr []byte)
	Hash       func(from, to int) int
}

func NewRollingHash() *RollingHash {
	const MASK30 = (1 << 30) - 1
	const MASK31 = (1 << 31) - 1
	const MOD = (1 << 61) - 1
	const POSITIVIZER = MOD * 4
	base := uint(2 + rand.Intn(MOD-4))

	mod := func(x uint) uint {
		xu := x >> 61
		xd := x & MOD
		ret := xu + xd
		if ret >= MOD {
			ret -= MOD
		}
		return ret
	}
	mul := func(a, b uint) uint {
		au := a >> 31
		ad := a & MASK31
		bu := b >> 31
		bd := b & MASK31
		mid := ad*bu + au*bd
		midu := mid >> 30
		midd := mid & MASK30
		return mod(au*bu*2 + midu + (midd << 31) + ad*bd)
	}

	var hash []uint
	var powmemo []uint

	initInts := func(arr []int) {
		hash = make([]uint, len(arr)+1)
		powmemo = make([]uint, len(arr)+1)
		powmemo[0] = 1
		for i, b := range arr {
			hash[i+1] = mod(mul(hash[i], base) + uint(b))
			powmemo[i+1] = mul(powmemo[i], base)
		}
	}

	initString := func(value string) {
		ret := make([]int, len(value))
		for i, v := range value {
			ret[i] = int(v)
		}
		initInts(ret)
	}

	initBytes := func(arr []byte) {
		ret := make([]int, len(arr))
		for i, v := range arr {
			ret[i] = int(v)
		}
		initInts(ret)
	}

	calcHash := func(from, to int) int {
		h := mod(hash[to] + POSITIVIZER - mul(hash[from], powmemo[to-from]))
		return int(h)
	}

	return &RollingHash{
		InitInts:   initInts,
		InitString: initString,
		InitBytes:  initBytes,
		Hash:       calcHash,
	}
}
