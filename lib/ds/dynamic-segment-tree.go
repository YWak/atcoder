package ds

type DynamicSegmentTree[V any] struct {
	// Initは[0, n)のsegment treeを初期化します。
	// 各要素の値は単位元となります。
	Init func(n int) *DynamicSegmentTree[V]

	// Getはi(0-based)番目の値を返します。
	Get func(i int) *V

	// Updateはi(0-based)番目の値をvalueに更新します。
	Update func(i int, value *V)

	// Queryは[l, r) (0-based)の計算値を返します。
	Query func(l, r int) *V
}

func NewDynamicSegmentTree[V any](
	e func() *V,
	calc func(a, b, ab *V),
) *DynamicSegmentTree[V] {
	st := DynamicSegmentTree[V]{}
	nodes := map[uint64]*V{}
	var size uint64

	st.Init = func(n int) *DynamicSegmentTree[V] {
		x := uint64(1)
		for x/2 < uint64(n+1) {
			x *= 2
		}
		size = x / 2

		return &st
	}
	st.Get = func(i int) *V {
		t := uint64(i) + size
		if node, ex := nodes[t]; ex {
			return node
		} else {
			node := e()
			nodes[t] = node
			return node
		}
	}
	st.Update = func(i int, value *V) {
		t := uint64(i) + size
		nodes[t] = value

		for {
			t /= 2
			if t == 0 {
				break
			}

			l, r := t*2+0, t*2+1
			node_l := nodes[l]
			node_r := nodes[r]

			// 計算しなくていいならスキップ
			if node_l == nil && node_r == nil {
				continue
			}

			if node_l == nil {
				nd := e()
				node_l = nd
			}
			if node_r == nil {
				nd := e()
				node_r = nd
			}

			node_lr := nodes[t]
			if node_lr == nil {
				nd := e()
				node_lr = nd
				nodes[t] = node_lr
			}
			calc(node_l, node_r, node_lr)
		}
	}
	st.Query = func(l, r int) *V {
		ret := e()
		for ll, rr := uint64(l)+size, uint64(r)+size; ll < rr; ll, rr = ll/2, rr/2 {
			var nx *V
			if ll%2 == 1 {
				l := nodes[ll]
				if l == nil {
					nx = ret
				} else {
					nx = e()
					calc(ret, l, nx)
				}
				ll++
			}
			if rr%2 == 1 {
				rr--
				r := nodes[rr]
				if r == nil {
					nx = ret
				} else {
					nx = e()
					calc(r, ret, nx)
				}
			}
			ret = nx
		}

		return ret
	}

	return &st
}
