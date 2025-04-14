package math

// Primes はn以下の素数を列挙する
func Primes(n int) []int {
	used := make([]bool, n+1)
	ps := []int{}

	for i := 2; i <= n; i++ {
		if used[i] {
			continue
		}
		ps = append(ps, i)
		for j := i; j <= n; j += i {
			used[j] = true
		}
	}
	return ps
}
