package algorithm

type ZAlgorithm struct{}

func (z ZAlgorithm) FromInts(values []int) []int {
	n := len(values)
	ret := make([]int, n)
	ret[0] = n
	i, j := 1, 0

	for i < n {
		for i+j < n && values[j] == values[i+j] {
			j++
		}
		ret[i] = j

		if j == 0 {
			i++
			continue
		}

		k := 1
		for i+k < n && k+ret[k] < j {
			ret[i+k] = ret[k]
			k++
		}
		i += k
		j -= k
	}

	return ret
}

func (z ZAlgorithm) FromString(value string) []int {
	s := make([]int, len(value))
	for i, v := range value {
		s[i] = int(v)
	}
	return z.FromInts(s)
}

func (z ZAlgorithm) FromBytes(values []byte) []int {
	a := make([]int, len(values))
	for i, v := range values {
		a[i] = int(v)
	}
	return z.FromInts(a)
}
