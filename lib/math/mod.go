package math

import "fmt"

type ModInt int

const Mod998244353 = ModInt(998244353)
const Mod1000000007 = ModInt(1000000007)

// Normはaをmod mの値に変換します
func (mod ModInt) Norm(a int) int {
	a %= int(mod)
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
	return mod.Norm((a * b) % int(mod))
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

// Fracは aをそれっぽい分数で表現します。
func (mod ModInt) Frac(a int) (int, int) {
	v1, v2 := int(mod), 0
	w1, w2 := a, 1

	for w1*w1*2 > int(mod) {
		q := v1 / w1
		z1, z2 := v1-q*w1, v2-q*w2

		v1, v2, w1, w2 = w1, w2, z1, z2
	}
	if w2 < 0 {
		w1, w2 = -w1, -w2
	}
	return w1, w2
}

func (mod ModInt) Fracs(a int) string {
	w1, w2 := mod.Frac(a)
	if w2 == 1 {
		return fmt.Sprintf("%d", w1)
	} else {
		return fmt.Sprintf("%d/%d", w1, w2)
	}
}

type Combination interface {
	// Factはnの階乗を返します。
	Fact(n int) int

	// Permはn個のモノからr個取り出して並べる順列の個数を返します。
	Perm(n, k int) int

	// Chooseはn個のモノからr個取り出す組み合わせの個数を返します。
	Choose(n, k int) int

	// RepChooseはn個のモノから重複を許して取り出す組み合わせの個数を返します。
	RepChoose(n, k int) int
}

type factCombination struct {
	mod   ModInt
	fact  []int
	ifact []int
}

func (com *factCombination) init(n int) {
	mod := com.mod
	com.fact = make([]int, n+1)
	com.ifact = make([]int, n+1)

	com.fact[1] = 1
	for i := 2; i <= n; i++ {
		com.fact[i] = mod.Mul(com.fact[i-1], i)
	}

	com.ifact[n] = mod.Inv(com.fact[n])
	for i := n - 1; i > 0; i-- {
		com.ifact[i] = mod.Mul(com.ifact[i+1], i)
	}
}

func (com *factCombination) Fact(n int) int {
	return com.fact[n]
}

func (com *factCombination) Perm(n, r int) int {
	if n <= 0 || r <= 0 || n < r {
		return 1
	}
	return com.mod.Mul(com.fact[n], com.ifact[n-r])
}

func (com *factCombination) Choose(n, r int) int {
	if n <= 0 || r <= 0 || n < r {
		return 1
	}
	return com.mod.Mul(com.fact[n], com.mod.Mul(com.ifact[r], com.ifact[n-r]))
}

func (com *factCombination) RepChoose(n, r int) int {
	return com.Choose(n+r-1, r)
}

var _ Combination = &factCombination{}

const MAX_FACT_SIZE = int(1e7)

// 組合せの計算を行うモジュールを返します。
func (mod ModInt) NewComination(nmax, kmax int) Combination {
	if nmax > MAX_FACT_SIZE || kmax > MAX_FACT_SIZE {
		panic(fmt.Sprintf("nmax and kmax must not be greater than %d", MAX_FACT_SIZE))
	}

	c := factCombination{mod: mod}
	c.init(Max(nmax, kmax))
	return &c
}
