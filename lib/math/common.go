package math

// INF18 は最大値を表す数
const INF18 = int(2e18) + int(2e9)

// INF9 は最大値を表す数
const INF9 = int(2e9)

// N10_6は10^6
const N10_6 = int(1e6)

// Maxはaとbのうち大きい方を返します。
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Minはaとbのうち小さい方を返します。
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Chmaxはaとbのうち大きい方をaに設定します。
func Chmax(a *int, b int) {
	*a = Max(*a, b)
}

// Chminはaとbのうち小さい方をaに設定します。
func Chmin(a *int, b int) {
	*a = Min(*a, b)
}

// Absはaの絶対値を返します。
func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// Powはaのb乗を返します。
func Pow(a, b int) int {
	ans := 1
	for b > 0 {
		if b%2 == 1 {
			ans *= a
		}
		a, b = a*a, b/2
	}
	return ans
}

// Divceilはa/b の結果を正の無限大に近づけるように丸めて返します。
func Divceil(a, b int) int {
	if a%b == 0 || a/b < 0 {
		return a / b
	}
	return (a + b - 1) / b
}

// Divfloorはa/bの結果を負の無限大に近づけるように丸めて返します。
func Divfloor(a, b int) int {
	if a%b == 0 || a/b > 0 {
		return a / b
	}
	if b < 0 {
		a, b = -a, -b
	}
	return (a - b + 1) / b
}

// Powmodは(x^n) mod m を返します。
func Powmod(x, n, m int) int {
	if n == 0 {
		return 1
	}

	x = x % m
	if x == 0 {
		return 0
	}

	ans := 1
	for n > 0 {
		if n%2 == 1 {
			ans = (ans * x) % m
		}
		x = (x * x) % m
		n /= 2
	}
	return ans
}

func Gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	aa, bb := a, b
	if bb > aa {
		aa, bb = bb, aa
	}
	for bb != 0 {
		aa, bb = bb, aa%bb
	}
	return a / aa * b
}
