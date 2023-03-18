package fgeo

import "math"

// FPointは浮動小数点表現の点をあらわします。
type FPoint struct {
	X float64
	Y float64
}

// FLineは直線を ax+by+cの形式で保存します。
type FLine struct {
	A float64
	B float64
	C float64
}

// FCircleは円をあらわします。
type FCircle struct {
	Center *FPoint
	Radius float64
}

func (a *FPoint) Add(b *FPoint) *FPoint {
	return &FPoint{a.X + b.X, a.Y + b.Y}
}

func (a *FPoint) Sub(b *FPoint) *FPoint {
	return &FPoint{a.X - b.X, a.Y - b.Y}
}

// RotateDegは原点を中心にして点aをdeg度回転させた点を返します。
func (a *FPoint) RotateDeg(deg float64) *FPoint {
	return a.RotateRad(deg * math.Pi / 180)
}

// RotateRadは原点を中心にして点aをth回転させた点を返します。
// thの単位はラジアンです。
func (a *FPoint) RotateRad(th float64) *FPoint {
	cos, sin := math.Cos(th), math.Sin(th)
	return &FPoint{a.X*cos - a.Y*sin, a.X*sin + a.Y*cos}
}

// Distは点aから点bまでの距離を返します。
func (a *FPoint) Dist(b *FPoint) float64 {
	x := a.X - b.X
	y := a.Y - b.Y
	return math.Sqrt(x*x + y*y)
}

// NewFLineFromPointsは点aと点bを通る直線を返します。
func NewFLineFromPoints(a, b *FPoint) *FLine {
	if a.X == b.X {
		if a.Y == b.Y {
			return nil
		}
		// x = X
		return &FLine{A: 1, B: 0, C: -a.X}
	}

	// 二点を通る直線の式からもとめる
	// y-y1 = (y1-y2)/(x1-x2) x - x1 => (y2-y1)x + (x1-x2)y + (x1-y1)(x1-x2) = 0
	return &FLine{A: b.Y - a.Y, B: a.X - b.X, C: (a.X - a.Y) * (a.X - b.X)}
}

// PerpendicularBisectorは点aと点bの垂直二等分線を返します。
func (a *FPoint) PerpendicularBisector(b *FPoint) *FLine {
	if a.Y == b.Y {
		if a.X == b.X {
			return nil
		}
		return &FLine{A: 1, B: 0, C: -(a.X + b.X) / 2}
	}

	c := &FPoint{X: (a.X + b.X) / 2, Y: (a.Y + b.Y) / 2}
	m := (b.X - a.X) / (a.Y - b.Y)
	return &FLine{A: -m, B: 1, C: m*c.X - c.Y}
}

// Crossingはl1とl2の交点を返します。
// https://mathwords.net/nityokusenkoten#axbyc0
func (l1 *FLine) Crossing(l2 *FLine) *FPoint {
	// 平行線は交わらない
	if l1.A*l2.B == l1.B*l2.A {
		return nil
	}

	return &FPoint{
		X: (l1.B*l2.C - l2.B*l1.C) / (l1.A*l2.B - l2.A*l1.B),
		Y: (l2.A*l1.C - l1.A*l2.C) / (l1.A*l2.B - l2.A*l1.B),
	}
}

func (c *FCircle) Includes(p *FPoint) bool {
	// 中心との距離がr以内である
	return c.Center.Dist(p) <= c.Radius
}

func (c *FCircle) From(ps []*FPoint) {
	switch len(ps) {
	case 0:
		c.Center = &FPoint{0.0, 0.0}
		c.Radius = 0.0
	case 1:
		c.Center = &FPoint{ps[0].X, ps[0].Y}
		c.Radius = 0.0
	case 2:
		t := NewCircumscribedCircle2(ps[0], ps[1])
		c.Center, c.Radius = t.Center, t.Radius
	default:
		t := NewCircumscribedCircle3(ps[0], ps[1], ps[2])
		c.Center, c.Radius = t.Center, t.Radius
	}
}

// 2点を半径とする円を返します。
func NewCircumscribedCircle2(a, b *FPoint) *FCircle {
	c := &FPoint{
		X: (a.X + b.X) / 2,
		Y: (a.Y + b.Y) / 2,
	}

	return &FCircle{Center: c, Radius: a.Dist(b) / 2}
}

// 3点の外接円を返します。
func NewCircumscribedCircle3(a, b, c *FPoint) *FCircle {
	// 垂直二等分線の交点
	l1 := a.PerpendicularBisector(b)
	l2 := b.PerpendicularBisector(c)

	o := l1.Crossing(l2)
	if o == nil {
		return nil
	}
	return &FCircle{Center: o, Radius: o.Dist(a)}
}

// SmallestEnclosingCircleはp1,p2を通る半径rの円の中心を返します。
func SmallestEnclosingCircle(points []*FPoint) *FCircle {
	// とりあえず大きいサイズ
	var circle *FCircle
	ok := func(c *FCircle) bool {
		for _, p := range points {
			if !c.Includes(p) {
				return false
			}
		}
		return true
	}

	// 2点を半径とする円
	for i, p := range points {
		for j, q := range points {
			if i == j {
				continue
			}
			c := NewCircumscribedCircle2(p, q)

			// 今より大きい円は無視する
			if (circle == nil || circle.Radius > c.Radius) && ok(c) {
				circle = c
			}
		}
	}

	//3点を通る円
	for i, p := range points {
		for j, q := range points {
			for k, r := range points {
				if i == j || j == k {
					continue
				}
				c := NewCircumscribedCircle3(p, q, r)
				if c != nil && (circle == nil || circle.Radius > c.Radius) && ok(c) {
					circle = c
				}
			}
		}
	}
	return circle
}
