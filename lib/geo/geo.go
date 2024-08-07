package geo

import (
	gomath "math"
	"sort"
)

type Point struct {
	x int
	y int
}

type Geo struct {
}

// NextPointは入力から2点読み込んで座標として返します。
// tested:
//
//	https://atcoder.jp/contests/abc250/tasks/abc250_f
func (g *Geo) NextPoint(a, b int) *Point {
	return &Point{a, b}
}

func (g *Geo) NextFPoint(x, y float64) *FPoint {
	return &FPoint{x, y}
}

// AreaX2は3点で囲まれる面積の2倍を返します。
// tested:
//
//	https://atcoder.jp/contests/abc250/tasks/abc250_f
func (g *Geo) AreaX2(p1 *Point, p2 *Point, p3 *Point) int {
	return (p1.x*p2.y + p2.x*p3.y + p3.x*p1.y - p1.y*p2.x - p2.y*p3.x - p3.y*p1.x)
}

// Dist2は2点間の距離の2乗を返します。
// tested:
//
//	https://atcoder.jp/contests/abc174/tasks/abc174_b
func (g *Geo) Dist2(p1 *Point, p2 *Point) int {
	x, y := p1.x-p2.x, p1.y-p2.y
	return x*x + y*y
}

// Distは2点間の距離を返します。
// tested:
//
//	https://atcoder.jp/contests/abc010/tasks/abc010_3
func (g *Geo) Dist(p1, p2 *Point) float64 {
	return gomath.Sqrt(float64(g.Dist2(p1, p2)))
}

// HenkakuSortは偏角によってソートを行います。
// tested:
//
//	https://atcoder.jp/contests/abc139/tasks/abc139_f
func (g *Geo) HenkakuSort(points []*Point) {
	sig := func(x, y int) int {
		if y > 0 {
			return 0
		}
		if y == 0 && x > 0 {
			return 0
		}
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		sp, sq := sig(p.x, p.y), sig(q.x, q.y)
		if sp != sq {
			return sp < sq
		}
		return p.x*q.y < p.y*q.x
	})
}
