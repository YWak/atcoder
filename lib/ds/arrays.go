package ds

func NewInt2d(a, b, value int) [][]int {
	arr := make([][]int, a)
	for i := range arr {
		arr[i] = make([]int, b)

		for j := range arr[i] {
			arr[i][j] = value
		}
	}

	return arr
}

func NewInt3d(a, b, c, value int) [][][]int {
	arr := make([][][]int, a)
	for i := range arr {
		arr[i] = NewInt2d(b, c, value)
	}
	return arr
}
