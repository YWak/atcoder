package ds

// 重み付きUnion-Findは集合と、各ノードがルートからどれくらい距離があるかを管理します。
type WeightedUnionFind struct {
	par  []int
	rank []int
	diff []int
}

// NewWeightedUnionFindは重み付きUnion-Findの実装を返します。
func NewWeightedUnionFind(n int) *WeightedUnionFind {
	uf := WeightedUnionFind{
		par:  make([]int, n),
		rank: make([]int, n),
		diff: make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.par[i] = i
		uf.rank[i] = 0
		uf.diff[i] = 0
	}

	return &uf
}

// xが接続されているルートを返します。
func (uf *WeightedUnionFind) Root(x int) int {
	if uf.par[x] == x {
		return x
	}
	r := uf.Root(uf.par[x])
	uf.diff[x] += uf.diff[uf.par[x]]
	uf.par[x] = r
	return r
}

// xの重みを返します。
func (uf *WeightedUnionFind) Weight(x int) int {
	uf.Root(x)
	return uf.diff[x]
}

// xとyが同じ集合にいるかを判断します。
func (uf *WeightedUnionFind) IsSame(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

// weight[y]-weight[x] = wとなるようにマージします
func (uf *WeightedUnionFind) Merge(x, y, w int) bool {
	w += uf.Weight(x)
	w -= uf.Weight(y)
	x = uf.Root(x)
	y = uf.Root(y)

	if x == y {
		return false
	}
	if uf.rank[x] < uf.rank[y] {
		w = -w
		x, y = y, x
	}
	if uf.rank[x] == uf.rank[y] {
		uf.rank[x]++
	}
	uf.par[y] = x
	uf.diff[y] = w
	return true
}

// weight[y]-weight[x]を返します。
func (uf *WeightedUnionFind) Diff(x, y int) int {
	return uf.Weight(y) - uf.Weight(x)
}
