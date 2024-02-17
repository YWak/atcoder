package ds

type LazySegmentTree[V, F any] struct {
	// Initは長さnの配列として初期化します
	Init func(n int) *LazySegmentTree[V, F]

	// InitByArrayはarrとして初期化します。
	InitByArray func(arr []V) *LazySegmentTree[V, F]

	// Updateは[l, r)をfで更新します。
	Update func(l, r int, f F)

	// Queryは[l, r)の値を返します。
	Query func(l, r int) V
}

// NewLazySegmentTreeは遅延評価セグメントツリーの実装を返します。
// Vは扱うモノイドの型、Fはモノイドに適用する写像をあらわします。
//
// 例えば range add range max queryの場合、以下のようになります。
//
// st := NewLazySegmentTree[int, int](
//
//	max,
//	func(f, x int) int { return f + x },
//	func(f, g int) int { return f + g },
//	func() int { return 0 },
//	func() int { return 0 },
//
// )
func NewLazySegmentTree[V, F any](
	operate func(a, b, ab *V),
	mapping func(f *F, x *V),
	composition func(f, g, fg *F),
	e func() V,
	id func() F,
) *LazySegmentTree[V, F] {
	st := &LazySegmentTree[V, F]{}
	var log int
	var size int
	var data []V
	var lazy []F

	update := func(k int) {
		operate(&data[k*2+0], &data[k*2+1], &data[k])
	}

	// 初期化
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.Init = func(n int) *LazySegmentTree[V, F] {
		arr := make([]V, n)
		for i := 0; i < n; i++ {
			arr[i] = e()
		}
		return st.InitByArray(arr)
	}

	// 配列による初期化
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.InitByArray = func(arr []V) *LazySegmentTree[V, F] {
		n := len(arr)
		size = 1
		log = 0
		for size < n {
			size *= 2
			log++
		}
		data = make([]V, size*2)
		for i := range data {
			data[i] = e()
		}
		for i := 0; i < n; i++ {
			data[size+i] = arr[i]
		}
		for i := size - 1; i >= 1; i-- {
			update(i)
		}
		lazy = make([]F, size)
		for i := range lazy {
			lazy[i] = id()
		}

		return st
	}

	applyAll := func(k int, f *F) {
		mapping(f, &data[k])
		if k < size {
			composition(f, &lazy[k], &lazy[k])
		}
	}
	push := func(k int) {
		applyAll(2*k+0, &lazy[k])
		applyAll(2*k+1, &lazy[k])
		lazy[k] = id()
	}
	// [l, r)の値を取得する。
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.Query = func(l, r int) V {
		if l == r {
			return e()
		}
		l, r = l+size, r+size
		for i := log; i >= 1; i-- {
			if (l>>i)<<i != l {
				push(l >> i)
			}
			if (r>>i)<<i != r {
				push(r >> i)
			}
		}
		sml, smr := e(), e()
		for l < r {
			if l%2 == 1 {
				operate(&sml, &data[l], &sml)
				l++
			}
			if r%2 == 1 {
				r--
				operate(&data[r], &smr, &smr)
			}
			l >>= 1
			r >>= 1
		}

		var ans V
		operate(&sml, &smr, &ans)
		return ans
	}

	// [l, r)の値をfで更新する
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.Update = func(l, r int, f F) {
		if l == r {
			return
		}
		l, r = l+size, r+size
		for i := log; i >= 1; i-- {
			if (l>>i)<<i != l {
				push(l >> i)
			}
			if (r>>i)<<i != r {
				push((r - 1) >> i)
			}
		}
		ll, rr := l, r
		for ll < rr {
			if ll%2 == 1 {
				applyAll(ll, &f)
				ll++
			}
			if rr%2 == 1 {
				rr--
				applyAll(rr, &f)
			}
			ll >>= 1
			rr >>= 1
		}
		for i := 1; i <= log; i++ {
			if (l>>i)<<i != l {
				update(l >> i)
			}
			if (r>>i)<<i != r {
				update((r - 1) >> i)
			}
		}
	}

	return st
}
