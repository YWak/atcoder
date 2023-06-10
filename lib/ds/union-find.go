package ds

// UnionFind : UnionFind構造を保持する構造体
type UnionFind struct {
	par []int // i番目のノードに対応する親
}

// NewUnionFindは、[0, n)のノードを持つUnion-Findを作る
func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{par: make([]int, n)}

	for i := 0; i < n; i++ {
		uf.par[i] = -1
	}

	return &uf
}

// Root はxのルートを得る
func (uf *UnionFind) Root(x int) int {
	if uf.par[x] < 0 {
		return x
	}
	uf.par[x] = uf.Root(uf.par[x])
	return uf.par[x]
}

// Unite はxとyを併合する。集合の構造が変更された(== 呼び出し前は異なる集合だった)かどうかを返す
func (uf *UnionFind) Unite(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)

	if rx == ry {
		return false
	}
	if uf.par[rx] > uf.par[ry] {
		rx, ry = ry, rx
	}
	uf.par[rx] += uf.par[ry]
	uf.par[ry] = rx
	return true
}

// Same はxとyが同じノードにいるかを判断する
func (uf *UnionFind) Same(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)
	return rx == ry
}

// Size は xの集合のサイズを返します。
func (uf *UnionFind) Size(x int) int {
	return -uf.par[uf.Root(x)]
}
