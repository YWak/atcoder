package utils

func bitOf(n, pos int) int {
	return (n >> uint(pos)) & 1
}

func bitOf64(n int64, pos int64) int {
	return int((n >> uint64(pos)) & 1)
}
