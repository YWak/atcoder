package utils

// for int --------------------------------
func max(a int, b ...int) int {
	l := len(b)

	for i := 0; i < l; i++ {
		if a < b[i] {
			a = b[i]
		}
	}

	return a
}

func min(a int, b ...int) int {
	l := len(b)

	for i := 0; i < l; i++ {
		if a > b[i] {
			a = b[i]
		}
	}

	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return +a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// for int --------------------------------
