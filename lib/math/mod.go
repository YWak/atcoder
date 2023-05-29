package math

type ModInt int

const Mod998244353 = ModInt(998244353)
const Mod1000000007 = ModInt(1000000007)

// Normはaをmod mの値に変換します
func (mod ModInt) Norm(a int) int {
	if a < 0 || a >= int(mod) {
		a %= int(mod)
	}
	if a < 0 {
		a += int(mod)
	}
	return a
}

// Addは a+b (mod m)を返します。
func (mod ModInt) Add(a, b int) int {
	ab := a + b
	if ab >= int(mod) {
		ab %= int(mod)
	}
	return ab
}

// Subは a-b (mod m)を返します。
func (mod ModInt) Sub(a, b int) int {
	ab := a - b
	if ab < 0 {
		ab += int(mod)
	}
	return ab
}

// Mulは a*b (mod m)を返します。
func (mod ModInt) Mul(a, b int) int {
	return (a * b) % int(mod)
}

// Powは(x^n) mod m を返します。
func (mod ModInt) Pow(x, n int) int {
	if n == 0 {
		return 1
	}

	x = x % int(mod)
	if x == 0 {
		return 0
	}

	ans := 1
	for n > 0 {
		if n%2 == 1 {
			ans = (ans * x) % int(mod)
		}
		x = (x * x) % int(mod)
		n /= 2
	}

	return ans
}

// Invはmod mにおけるaの逆元を返します。
func (mod ModInt) Inv(a int) int {
	// 拡張ユークリッドの互除法
	b, u, v := int(mod), 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	return mod.Norm(u)
}

// Divはa / b (mod m)を返します。
func (mod ModInt) Div(a, b int) int {
	return mod.Mul(a, mod.Inv(b))
}

// Chaddはa + b (mod m)の結果をaに設定します。
func (mod ModInt) Chadd(a *int, b int) {
	*a = mod.Add(*a, b)
}

// Chsubはa - b (mod m)の結果をaに設定します。
func (mod ModInt) Chsub(a *int, b int) {
	*a = mod.Sub(*a, b)
}

// Chmulはa * b (mod m)の結果をaに設定します。
func (mod ModInt) Chmul(a *int, b int) {
	*a = mod.Mul(*a, b)
}

// Chdivはa / b (mod m)の結果をaに設定します。
func (mod ModInt) Chdiv(a *int, b int) {
	*a = mod.Div(*a, b)
}
