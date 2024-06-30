package ds

import (
	"github.com/ywak/atcoder/lib/math"
)

// Graph はグラフを表現する構造です
type Graph struct {
	// 隣接リスト
	List [][]*Edge
	// 辺のリスト
	Edges []*Edge
}

// Edge は辺を表現する構造体です
type Edge struct {
	From int

	// 行き先
	To int

	// 重み
	Weight int

	// 辺の番号 (0-indexed)
	Index int
}

// NewGraph はグラフを作成します
func NewGraph(n int) *Graph {
	return &Graph{
		make([][]*Edge, n),
		make([]*Edge, 0),
	}
}

// CountNodesはノードの数を返します。
func (g *Graph) CountNodes() int {
	return len(g.List)
}

// AddEdge uからvへの重み1の無向辺を追加します。
func (g *Graph) AddEdge(u, v int) {
	g.AddWeightedEdge(u, v, 1)
}

// AddEdge1はuからv(1-indexed)への重み1の無向辺を追加します。
func (g *Graph) AddEdge1(u, v int) {
	g.AddEdge(u-1, v-1)
}

// AddWeightedEdge はuからvへ重みwの無向辺を追加します。
func (g *Graph) AddWeightedEdge(u, v, w int) {
	i := len(g.Edges)
	e := Edge{u, v, w, i}
	g.List[u] = append(g.List[u], &e)
	g.List[v] = append(g.List[v], &Edge{v, u, w, i})

	g.Edges = append(g.Edges, &e)
}

// AddDirectedEdgeはuからvへ重み1の有向辺を追加します。
func (g *Graph) AddDirectedEdge(u, v int) {
	g.AddDirectedWeightedEdge(u, v, 1)
}

// AddDirectedWeightedEdgeはuからvへ重みwの有向辺を追加します。
func (g *Graph) AddDirectedWeightedEdge(u, v, w int) {
	i := len(g.Edges)
	e := Edge{u, v, w, i}
	g.List[u] = append(g.List[u], &e)
	g.Edges = append(g.Edges, &e)
}

type DijkstraResult struct {
	// Costsは各頂点へのコスト
	Costs []int

	// Prevsは各頂点に最短距離で到着したとき、直前にいた頂点。開始地点は-1になる
	Prevs []int
}

// Pathはtまでの頂点を順に返します。
func (dr *DijkstraResult) Path(t int) []int {
	path := []int{}
	for p := t; p != -1; p = dr.Prevs[p] {
		path = append(path, p)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// Dijkstra はsから各頂点への最短距離を返します。
// 重みが負の辺があるときには使用できません。
// 計算量: |V| + |E|log|V|
func (g *Graph) Dijkstra(s int) *DijkstraResult {

	n := g.CountNodes()

	// 初期化
	cost := make([]int, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.INF18
	}
	cost[s] = 0
	prev[s] = -1

	type Node struct{ node, cost int }
	pq := NewPriorityQueue(func(a, b *Node) bool {
		return a.cost < b.cost
	})
	pq.Push(&Node{s, 0})

	for pq.HasElements() {
		u := pq.Pop()

		if cost[u.node] < u.cost {
			// 更新されていたら無視する
			continue
		}

		for _, e := range g.List[u.node] {
			c := cost[u.node] + e.Weight

			if cost[e.To] > c {
				cost[e.To] = c
				prev[e.To] = u.node
				pq.Push(&Node{e.To, c})
			}
		}
	}

	return &DijkstraResult{Costs: cost, Prevs: prev}
}

// BellmanFord はsからtへの最短ルートを返します。
func (g *Graph) BellmanFord(s, t int) {
}

// WarshallFloyd は全点対の最短ルートを返します。
func (g *Graph) WarshallFloyd() [][]int {
	n := g.CountNodes()
	d := NewInt2d(n, n, math.INF18)
	for i := 0; i < n; i++ {
		d[i][i] = 0
		for _, e := range g.List[i] {
			d[i][e.To] = e.Weight
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = math.Min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	return d
}

// MaxFlowはsからtへの最大流を求めます。
func (g *Graph) MaxFlow(s, t int) int {
	type DinicEdge struct {
		To  int
		Rev int
		Cap int
	}

	_g := make([][]*DinicEdge, g.CountNodes())
	for u := range g.List {
		for _, e := range g.List[u] {
			le1 := len(_g[u])
			le2 := len(_g[e.To])

			_g[u] = append(_g[u], &DinicEdge{e.To, le2, e.Weight})
			_g[e.To] = append(_g[e.To], &DinicEdge{u, le1, 0})
		}
	}

	level := make([]int, len(_g))
	bfs := func() {
		for i := range level {
			level[i] = -1
		}
		level[s] = 0
		queue := []int{s}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			for _, e := range _g[p] {
				if e.Cap > 0 && level[e.To] < 0 {
					level[e.To] = level[p] + 1
					queue = append(queue, e.To)
				}
			}
		}
	}

	var loopstart []int
	var dfs func(curr, f int) int
	dfs = func(curr, f int) int {
		if curr == t {
			return f
		}
		for i := loopstart[curr]; i < len(_g[curr]); i++ {
			loopstart[curr] = i
			e := _g[curr][i]

			if e.Cap > 0 && level[curr] < level[e.To] {
				d := dfs(e.To, math.Min(f, e.Cap))
				if d > 0 {
					e.Cap -= d
					_g[e.To][e.Rev].Cap += d
					return d
				}
			}
		}
		return 0
	}

	flow := 0
	for {
		bfs()
		if level[t] < 0 {
			return flow
		}

		loopstart = make([]int, len(_g))
		for {
			f := dfs(s, math.INF18)
			if f == 0 {
				break
			}
			flow += f
		}
	}
}

// Sccは強連結成分分解の結果を表します。
type Scc struct {
	// nは強連結成分分解後の頂点数
	Size int
	// Graphは強連結成分分解後のグラフ
	Graph [][]int
	// Componentsは強連結成分分解後の頂点(0-indexed)と、その頂点に対応する、もとのグラフの頂点(0-indexed)の配列
	Components [][]int
	// Nodesはもとのグラフの頂点(0-indexed)と、対応する新しい頂点(0-indexed)
	Nodes map[int]int
}

// sccは与えられたグラフgを強連結成分分解した結果を返します。
func (g *Graph) Scc() Scc {
	s := Scc{
		Nodes: map[int]int{},
	}

	// 逆向きのグラフを作る
	rg := make([][]int, g.CountNodes())
	for u := range g.List {
		for _, e := range g.List[u] {
			rg[e.To] = append(rg[e.To], u)
		}
	}

	used := make([]bool, g.CountNodes())
	order := make([]int, 0, g.CountNodes())

	// 帰りがけ順に順序を保存するdfs
	var dfs func(u int)
	dfs = func(u int) {
		used[u] = true
		for _, e := range g.List[u] {
			if !used[e.To] {
				dfs(e.To)
			}
		}
		order = append(order, u)
	}
	for i := 0; i < g.CountNodes(); i++ {
		if !used[i] {
			dfs(i)
		}
	}

	// 逆向きにたどって採番する
	var rdfs func(u, k int)
	rdfs = func(u, k int) {
		s.Nodes[u] = k
		for _, v := range rg[u] {
			if _, e := s.Nodes[v]; !e {
				rdfs(v, k)
			}
		}
	}
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if _, e := s.Nodes[u]; !e {
			rdfs(u, s.Size)
			s.Size++
		}
	}

	// nodesをもとにしてcomponentsを作成する
	s.Components = make([][]int, s.Size)
	for from, to := range s.Nodes {
		s.Components[to] = append(s.Components[to], from)
	}

	type p struct {
		u, v int
	}

	// 強連結成分分解後のグラフを作成する
	connected := map[p]bool{}
	s.Graph = make([][]int, s.Size)
	for u, l := range g.List {
		for _, e := range l {
			if s.Nodes[u] != s.Nodes[e.To] && !connected[p{u, e.To}] {
				connected[p{u, e.To}] = true
				s.Graph[s.Nodes[u]] = append(s.Graph[s.Nodes[u]], s.Nodes[e.To])
			}
		}
	}

	return s
}

// Lcaは最小共通祖先を提供します。
type Lca struct {
	// nはグラフの要素数
	n int
	// xはグラフの階層の最大数
	x int
	// depthは要素ごとのルートからの深さ
	depth []int
	// ancestorsは要素ごとの2^k個上の祖先
	ancestors [][]int
}

// Lcaはrootを木の根とした最小共通祖先を作成します。
func (g *Graph) Lca(root int) *Lca {
	// 初期化
	n := g.CountNodes()
	x := 1
	for (1 << x) <= n {
		x++
	}

	lca := Lca{n, x, make([]int, n), NewInt2d(n, x, 0)}

	// depthとparentの初期化
	parent := make([]int, n)

	var dfs func(curr, prev, level int)
	dfs = func(curr int, prev int, level int) {
		lca.depth[curr] = level
		parent[curr] = prev

		for _, e := range g.List[curr] {
			if e.To == prev {
				continue
			}
			dfs(e.To, curr, level+1)
		}
	}
	dfs(root, -1, 0)

	// 2^k先の親を取得
	for k := 0; k < lca.x; k++ {
		for i := 0; i < lca.n; i++ {
			if k == 0 {
				lca.ancestors[i][k] = parent[i]
				continue
			}

			p := lca.ancestors[i][k-1]
			if p == -1 {
				lca.ancestors[i][k] = -1
			} else {
				lca.ancestors[i][k] = lca.ancestors[p][k-1]
			}
		}
		a := make([]int, lca.n)
		for i, v := range lca.ancestors {
			a[i] = v[k]
		}
	}

	return &lca
}

// Ofはaとbの最小共通祖先を返します。
func (lca *Lca) Of(a, b int) int {
	if lca.depth[a] > lca.depth[b] {
		a, b = b, a
	}
	// 高さを同じにする
	for k := lca.x - 1; k >= 0; k-- {
		if ((lca.depth[b]-lca.depth[a])>>k)%2 == 1 {
			b = lca.ancestors[b][k]
		}
	}
	if a == b {
		return a
	}

	// 同じ親にならないギリギリまで進める
	for k := lca.x - 1; k >= 0; k-- {
		if lca.ancestors[a][k] != lca.ancestors[b][k] {
			a = lca.ancestors[a][k]
			b = lca.ancestors[b][k]
		}
	}
	// 一個上が共通の親
	return lca.ancestors[a][0]
}

// Distanceはaとbの距離を返します。
func (lca *Lca) Distance(a, b int) int {
	c := lca.Of(a, b)
	return lca.depth[a] + lca.depth[b] - 2*lca.depth[c]
}
