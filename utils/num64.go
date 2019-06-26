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

// for int64 ------------------------------
