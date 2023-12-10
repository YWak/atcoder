package ds

import (
	math "github.com/ywak/atcoder/lib/math"
)

type SegmentTree struct {
	// このsegment treeが管理するインデックスの範囲。[0, n)を管理する。
	n int

	// segment treeの各ノードの値を保持する配列
	nodes []int

	e func() int

	calc func(a, b int) int
}

// NewSegmentTreeは区間和を扱うSegmentTreeを返します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func NewSegmentTree(
	e func() int,
	calc func(a, b int) int,
) *SegmentTree {
	return &SegmentTree{-1, []int{}, e, calc}
}

// NewRangeMaxQueryは区間最大値を扱うSegmentTreeを返します。
func NewRangeMaxQuery() *SegmentTree {
	return NewSegmentTree(func() int { return -math.INF18 }, func(a, b int) int { return math.Max(a, b) })
}

// NewRangeMinQueryは区間最小値を扱うSegmentTreeを返します。
// tested:
//
//	https://judge.yosupo.jp/problem/staticrmq
func NewRangeMinQuery() *SegmentTree {
	return NewSegmentTree(func() int { return math.INF18 }, func(a, b int) int { return math.Min(a, b) })
}

// Initは[0, n)のsegment treeを初期化します。
// 各要素の値は単位元となります。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func (st *SegmentTree) Init(n int) {
	// xはn*2を超える最小の2べき
	x := 1
	for x/2 < n+1 {
		x *= 2
	}
	st.n = x / 2
	st.nodes = make([]int, x)
	for i := 0; i < x; i++ {
		st.nodes[i] = st.e()
	}
}

// InitAsArrayはvalsで配列を初期化します。
// 区間の長さはlen(vals)になります。
// tested:
//
//	https://judge.yosupo.jp/problem/staticrmq
func (st *SegmentTree) InitAsArray(vals []int) {
	n := len(vals)
	// xはn*2を超える最小の2べき
	x := 1
	for x/2 < n {
		x *= 2
	}
	st.n = x / 2
	st.nodes = make([]int, x)

	for i, v := range vals {
		st.nodes[i+st.n] = v
	}
	for i := st.n - 1; i > 0; i-- {
		st.nodes[i] = st.calc(st.nodes[i*2], st.nodes[i*2+1])
	}
}

// Updateはi(0-based)番目の値をvalueに更新します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func (st *SegmentTree) Update(i int, value int) {
	t := i + st.n
	st.nodes[t] = value

	for {
		t /= 2
		if t == 0 {
			break
		}
		st.nodes[t] = st.calc(st.nodes[t*2], st.nodes[t*2+1])
	}
}

// Queryは[l, r) (0-based)の計算値を返します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func (st *SegmentTree) Query(l, r int) int {
	ret := st.e()
	for ll, rr := l+st.n, r+st.n; ll < rr; ll, rr = ll/2, rr/2 {
		if ll%2 == 1 {
			ret = st.calc(ret, st.nodes[ll])
			ll++
		}
		if rr%2 == 1 {
			rr--
			ret = st.calc(st.nodes[rr], ret)
		}
	}

	return ret
}

// Getはi番目(0-based)の要素を返します。
func (st *SegmentTree) Get(i int) int {
	return st.nodes[i+st.n]
}

// Allは全区間に対する値を返します。
func (st *SegmentTree) All() int {
	return st.nodes[1]
}
