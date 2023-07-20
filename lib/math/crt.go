package math

// crtは中国剰余定理の実装です。
// x = b[i] mod m[i]
// となる x mod lcm(m) を計算し、 x, lcm(m) を返します。
// 解なしの場合は 0, 0を返します。
func Crt(b, m []int) (int, int) {
	b0, m0 := 0, 1
	invgcd := func(a, m int) (int, int) {
		x, u := 1, 0
		for m != 0 {
			t := a / m
			a, m = m, a-t*m
			x, u = u, x-t*u
		}
		return a, x
	}
	for i := 0; i < len(b); i++ {
		b1, m1 := b[i], m[i]
		if m0 < m1 {
			b0, b1 = b1, b0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if b0%m1 != b1 {
				return 0, 0
			}
			continue
		}
		g, im := invgcd(m0, m1)
		if (b1-b0)%g != 0 {
			return 0, 0
		}
		u := m1 / g
		x := (b1 - b0) / g % u * im % u
		b0 += m0 * x
		m0 *= u
	}
	return (b0%m0 + m0) % m0, m0
}
