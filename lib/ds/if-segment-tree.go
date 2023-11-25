package ds

type IfSegmentTreeFunctions struct {
	// 単位元を返します
	E func() interface{}
	// 計算結果を返します
	Calc func(a, b interface{}) interface{}
}

type IfSegmentTree struct {
	// このsegment treeが管理するインデックスの範囲。[0, n)を管理する。
	n int

	// segment treeの各ノードの値を保持する配列
	nodes []interface{}

	// このsegment treeの値を操作する関数群
	F IfSegmentTreeFunctions
}

// NewSegmentTreeは区間和を扱うSegmentTreeを返します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func NewIfSegmentTree() *IfSegmentTree {
	return &IfSegmentTree{
		-1,
		[]interface{}{},
		IfSegmentTreeFunctions{
			func() interface{} { return []int{} },
			func(a, b interface{}) interface{} { return a },
		},
	}
}

// InitAsArrayはvalsで配列を初期化します。
// 区間の長さはlen(vals)になります。
// tested:
//
//	https://judge.yosupo.jp/problem/staticrmq
func (st *IfSegmentTree) InitAsArray(vals []interface{}) {
	n := len(vals)
	// xはn*2を超える最小の2べき
	x := 1
	for x/2 < n {
		x *= 2
	}
	st.n = x / 2
	st.nodes = make([]interface{}, x)

	for i, v := range vals {
		st.nodes[i+st.n] = v
	}
	for i := st.n - 1; i > 0; i-- {
		st.nodes[i] = st.F.Calc(st.nodes[i*2], st.nodes[i*2+1])
	}
}

// Updateはi(0-based)番目の値をvalueに更新します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func (st *IfSegmentTree) Update(i int, value interface{}) {
	t := i + st.n
	st.nodes[t] = value

	for {
		t /= 2
		if t == 0 {
			break
		}
		st.nodes[t] = st.F.Calc(st.nodes[t*2], st.nodes[t*2+1])
	}
}

// Queryは[l, r) (0-based)の計算値を返します。
// tested:
//
//	https://atcoder.jp/contests/abl/tasks/abl_d
func (st *IfSegmentTree) Query(l, r int) interface{} {
	ret := st.F.E()
	for ll, rr := l+st.n, r+st.n; ll < rr; ll, rr = ll/2, rr/2 {
		if ll%2 == 1 {
			ret = st.F.Calc(ret, st.nodes[ll])
			ll++
		}
		if rr%2 == 1 {
			rr--
			ret = st.F.Calc(st.nodes[rr], ret)
		}
	}

	return ret
}

// Getはi番目(0-based)の要素を返します。
func (st *IfSegmentTree) Get(i int) interface{} {
	return st.nodes[i+st.n]
}

// Allは全区間に対する値を返します。
func (st *IfSegmentTree) All() interface{} {
	return st.nodes[1]
}
