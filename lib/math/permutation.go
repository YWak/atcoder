package math

type Permutation []int

// Nextは次の順列を生成し、生成に成功したかどうかを返します。
func (p *Permutation) Next() bool {
	for i := len(*p) - 2; i >= 0; i-- {
		if (*p)[i] > (*p)[i+1] {
			continue
		}
		j := len(*p)
		for {
			j--
			if (*p)[i] < (*p)[j] {
				break
			}
		}
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
		for k, l := i+1, len(*p)-1; k < l; k, l = k+1, l-1 {
			(*p)[k], (*p)[l] = (*p)[l], (*p)[k]
		}
		return true

	}
	return false
}
