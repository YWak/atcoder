package math

type LinearSieve struct {
	// LPF[x]はxの最小素因数
	LPF []int

	// Primesは素数のリスト
	Primes []int
}

// Factorizeはxを素因数分解し、素因数とその指数のmapを返します。
func (ls *LinearSieve) Factorize(x int) map[int]int {
	pe := map[int]int{}
	for x != 1 {
		pe[ls.LPF[x]]++
		x /= ls.LPF[x]
	}

	return pe
}

// IsNthPowerは x = a^nとなる自然数aが存在するかを判断します。
func (ls *LinearSieve) IsNthPower(x, n int) bool {
	pe := ls.Factorize(x)

	for _, e := range pe {
		if e%n != 0 {
			return false
		}
	}
	return true
}

// NewLinearSieveはmaxまでの値を扱える線形篩を作成して返します。
func NewLinearSieve(max int) *LinearSieve {
	lpf := make([]int, max+1)
	if max > 1 {
		lpf[1] = 1
	}

	// 素因数のリスト
	ps := []int{}
	for v := 2; v <= max; v++ {
		if lpf[v] == 0 {
			lpf[v] = v
			ps = append(ps, v)
		}
		for _, p := range ps {
			if p*v > max || p > lpf[v] {
				break
			}
			lpf[p*v] = p
		}
	}

	return &LinearSieve{lpf, ps}
}
