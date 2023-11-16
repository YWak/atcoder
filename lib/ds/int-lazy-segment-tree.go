package ds

type IntLazySegmentTree struct {
	// Initは長さnの配列として初期化します
	Init func(n int) *IntLazySegmentTree

	// InitByArrayはarrとして初期化します。
	InitByArray func(arr []int) *IntLazySegmentTree

	// Updateは[l, r)をfで更新します。
	Update func(l, r int, f int)

	// Queryは[l, r)の値を返します。
	Query func(l, r int) int
}

// NewIntLazySegmentTreeは遅延評価セグメントツリーの実装を返します。
// Vは扱うモノイドの型、Fはモノイドに適用する写像をあらわします。
func NewIntLazySegmentTree(
	operate func(a, b int) int,
	mapping func(f int, x int) int,
	composition func(f, g int) int,
	e func() int,
	id func() int,
) *IntLazySegmentTree {
	st := &IntLazySegmentTree{}
	var log int
	var size int
	var data []int
	var lazy []int

	update := func(k int) {
		data[k] = operate(data[k*2+0], data[k*2+1])
	}

	// 初期化
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.Init = func(n int) *IntLazySegmentTree {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = e()
		}
		return st.InitByArray(arr)
	}

	// 配列による初期化
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.InitByArray = func(arr []int) *IntLazySegmentTree {
		n := len(arr)
		size = 1
		log = 0
		for size < n {
			size *= 2
			log++
		}
		data = make([]int, size*2)
		for i := 0; i < n; i++ {
			data[size+i] = arr[i]
		}
		for i := size - 1; i >= 1; i-- {
			update(i)
		}
		lazy = make([]int, size)
		for i := range lazy {
			lazy[i] = id()
		}

		return st
	}

	applyAll := func(k int, f int) {
		data[k] = mapping(f, data[k])
		if k < size {
			lazy[k] = composition(f, lazy[k])
		}
	}
	push := func(k int) {
		applyAll(2*k+0, lazy[k])
		applyAll(2*k+1, lazy[k])
		lazy[k] = id()
	}
	// [l, r)の値を取得する。
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.Query = func(l, r int) int {
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
				sml = operate(sml, data[l])
				l++
			}
			if r%2 == 1 {
				r--
				smr = operate(data[r], smr)
			}
			l >>= 1
			r >>= 1
		}

		return operate(sml, smr)
	}

	// [l, r)の値をfで更新する
	// tested:
	//   https://atcoder.jp/contests/typical90/tasks/typical90_ac
	st.Update = func(l, r int, f int) {
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
				applyAll(ll, f)
				ll++
			}
			if rr%2 == 1 {
				rr--
				applyAll(rr, f)
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
