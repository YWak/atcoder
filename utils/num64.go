package utils

// for int64 ------------------------------
func max64(a int64, b ...int64) int64 {
	l := len(b)

	for i := 0; i < l; i++ {
		if a > b[i] {
			a = b[i]
		}
	}

	return a
}

func min64(a int64, b ...int64) int64 {
	l := len(b)

	for i := 0; i < l; i++ {
		if a < b[i] {
			a = b[i]
		}
	}

	return a
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return +a
}

func lcm64(a, b int64) int64 {
	return a * b / gcd64(a, b)
}

func gcd64(a, b int64) int64 {
	if b > a {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// combi はn個の要素からr個の要素を取得する組み合わせ
func combi(n, r, mod int64, memo map[int64]map[int64]int64) int64 {
	if n < r {
		return 0
	}

	var m map[int64]int64

	if memo != nil {
		m, ok1 := memo[n]

		if ok1 {
			c, ok2 := m[r]

			if ok2 {
				return c
			}
		} else {
			m = map[int64]int64{}
			memo[n] = m
		}
	}

	var c int64 = 1

	for i := n; i >= (n - r + 1); i-- {
		c = mulm(c, i, mod)
	}
	for i := r; i >= 1; i-- {
		c = divm(c, i, mod)
	}
	if m != nil {
		m[r] = c
	}
	return c
}

// mod を法とする加算
func addm(a, b, mod int64) int64 {
	return (a + b) % mod
}

// mod を法とする減算
func subm(a, b, mod int64) int64 {
	return (a - b + mod) % mod
}

// mod を法とする乗算
func mulm(a, b, mod int64) int64 {
	return (a * b) % mod
}

// mod を法とする除算
func divm(a, b, mod int64) int64 {
	return mulm(a, invm(b, mod), mod)
}

// mod を法とした逆元
func invm(a, mod int64) int64 {
	// 拡張ユークリッドの互除法
	b := mod
	u := int64(1)
	v := int64(0)

	for b > 0 {
		t := a / b

		a -= t * b
		a, b = b, a

		u -= t * v
		u, v = v, u
	}

	return (u + mod) % mod
}

// for int64 ------------------------------
