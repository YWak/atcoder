package ds

// Vは扱う値、Xは保持する値、Fは適用する値
type DualSegmentTree[V, X, F any] struct {
	Init func(n int) *DualSegmentTree[V, X, F]

	Update func(l, r int, value F)

	Query func(i int) V
}

func NewDualSegmentTree[V, X, F any](
	// mappingは値vにxを適用した結果を保存します。
	mapping func(x *X, v *V),

	// compositionはxに関数fを適用した結果を保存します。
	composition func(f *F, x *X),

	// eは値の初期値を返します。
	e func() V,

	// idはXの初期値を返します。
	id func() X,
) *DualSegmentTree[V, X, F] {
	st := &DualSegmentTree[V, X, F]{}
	size := 1
	var nodes []X
	st.Init = func(n int) *DualSegmentTree[V, X, F] {
		log := 0
		for size < n {
			size *= 2
			log++
		}
		nodes = make([]X, size*2)
		for i := range nodes {
			nodes[i] = id()
		}

		return st
	}

	st.Query = func(i int) V {
		value := e()
		for t := i + size; t > 0; t /= 2 {
			mapping(&nodes[t], &value)
		}
		return value
	}

	st.Update = func(l, r int, value F) {
		for ll, rr := l+size, r+size; ll < rr; ll, rr = ll/2, rr/2 {
			if ll%2 == 1 {
				composition(&value, &nodes[ll])
				ll++
			}
			if rr%2 == 1 {
				rr--
				composition(&value, &nodes[rr])
			}
		}
	}

	return st
}
