package ds

type Int2d [][]int

func NewInt2d(a, b, value int) [][]int {
	base := make([]int, a*b)
	if value != 0 {
		for i := range base {
			base[i] = value
		}
	}
	arr := make([][]int, a)
	for i := range arr {
		arr[i] = base[(i * b):((i + 1) * b)]
	}

	return Int2d(arr)
}

func (arr *Int2d) Init(v int) {
	_a := (*arr)
	a := _a[0][0 : len(_a)*len(_a[0])]
	for i := 0; i < len(a); i++ {
		a[i] = v
	}
}

func NewInt3d(a, b, c, value int) [][][]int {
	arr := make([][][]int, a)
	for i := range arr {
		arr[i] = NewInt2d(b, c, value)
	}
	return arr
}

type Int4d [][][][]int

func NewInt4d(a, b, c, d, value int) Int4d {
	base := make([]int, a*b*c*d)
	if value != 0 {
		for i := range base {
			base[i] = value
		}
	}
	arr := make(Int4d, a)
	for i := range arr {
		arr[i] = make([][][]int, b)
		for j := range arr[i] {
			arr[i][j] = make([][]int, c)
			for k := range arr[i][j] {
				p := i*b*c*d + j*c*d + k*d
				arr[i][j][k] = base[p:(p + d)]
			}
		}
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
