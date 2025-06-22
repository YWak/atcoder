package ds

type SegmentTree[V any] struct {
	// Initは[0, n)のsegment treeを初期化します。
	// 各要素の値は単位元となります。
	// tested:
	//
	//	https://atcoder.jp/contests/abl/tasks/abl_d
	Init func(n int) *SegmentTree[V]

	// InitAsArrayはvalsで配列を初期化します。
	// 区間の長さはlen(vals)になります。
	// tested:
	//
	//	https://judge.yosupo.jp/problem/staticrmq
	InitAsArray func(arr []V) *SegmentTree[V]

	// Updateはi(0-based)番目の値をvalueに更新します。
	// tested:
	//
	//	https://atcoder.jp/contests/abl/tasks/abl_d
	Update func(i int, value V)

	// Queryは[l, r) (0-based)の計算値を返します。
	// tested:
	//
	//	https://atcoder.jp/contests/abl/tasks/abl_d
	Query func(l, r int) V

	// Getはi(0-based)番目の値を返します。
	Get func(i int) V
}

// NewSegmentTreeは区間和を扱うSegmentTreeを返します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func NewSegmentTree[V any](
	e func() V,
	calc func(a, b, ab *V),
) *SegmentTree[V] {
	st := SegmentTree[V]{}
	var nodes []V
	var size int
	init := func(n int) *SegmentTree[V] {
		// xはn*2を超える最小の2べき
		x := 1
		for x/2 < n+1 {
			x *= 2
		}
		size = x / 2
		nodes = make([]V, x)
		for i := 0; i < x; i++ {
			nodes[i] = e()
		}
		return &st
	}
	initAsArray := func(vals []V) *SegmentTree[V] {
		init(len(vals))

		for i, v := range vals {
			nodes[i+size] = v
		}
		for i := size - 1; i > 0; i-- {
			calc(&nodes[i*2], &nodes[i*2+1], &nodes[i])
		}

		return &st
	}
	st.Update = func(i int, value V) {
		t := i + size
		nodes[t] = value

		for {
			t /= 2
			if t == 0 {
				break
			}
			calc(&nodes[t*2], &nodes[t*2+1], &nodes[t])
		}
	}
	st.Query = func(l, r int) V {
		ret := e()
		for ll, rr := l+size, r+size; ll < rr; ll, rr = ll/2, rr/2 {
			if ll%2 == 1 {
				calc(&ret, &nodes[ll], &ret)
				ll++
			}
			if rr%2 == 1 {
				rr--
				calc(&nodes[rr], &ret, &ret)
			}
		}

		return ret
	}
	st.Get = func(i int) V {
		return nodes[i+size]
	}

	st.Init = init
	st.InitAsArray = initAsArray

	return &st
}

// // Getはi番目(0-based)の要素を返します。
// func (st *SegmentTree) Get(i int) int {
// 	return st.nodes[i+st.n]
// }

// // Allは全区間に対する値を返します。
// func (st *SegmentTree) All() int {
// 	return st.nodes[1]
// }
