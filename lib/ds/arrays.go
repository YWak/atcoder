package ds

func NewInt2d(a, b, value int) [][]int {
	arr := make([][]int, a)
	for i := range arr {
		arr[i] = make([]int, b)

		if value != 0 {
			for j := range arr[i] {
				arr[i][j] = value
			}
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

func NewBool2d(a, b int, value bool) [][]bool {
	arr := make([][]bool, a)
	for i := range arr {
		arr[i] = make([]bool, b)

		if value {
			for j := 0; j < b; j++ {
				arr[i][j] = true
			}
		}
	}

	return arr
}

func NewFloat2d(a, b int, value float64) [][]float64 {
	arr := make([][]float64, a)
	for i := range arr {
		arr[i] = make([]float64, b)

		if value != 0 {
			for j := range arr[i] {
				arr[i][j] = value
			}
		}
	}

	return arr
}
