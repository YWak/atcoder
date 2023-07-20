package math

type Permutation struct {
	values []int
}

func NewPermutation(n int) *Permutation {
	return NewPermutationFrom(0, n)
}

func NewPermutationFrom(s, n int) *Permutation {
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = s + i
	}

	return &Permutation{values: values}
}

func (p *Permutation) Get(i int) int {
	return p.values[i]
}

// Nextは次の順列を生成し、生成に成功したかどうかを返します。
func (p *Permutation) Next() bool {
	for i := len(p.values) - 2; i >= 0; i-- {
		if p.values[i] > p.values[i+1] {
			continue
		}
		j := len(p.values)
		for {
			j--
			if p.values[i] < p.values[j] {
				break
			}
		}
		p.values[i], p.values[j] = p.values[j], p.values[i]
		for k, l := i+1, len(p.values)-1; k < l; k, l = k+1, l-1 {
			p.values[k], p.values[l] = p.values[l], p.values[k]
		}
		return true
	}
	return false
}
