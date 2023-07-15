package math

import "sort"

// 座標圧縮を行います。mのkeyが昇順になるように0からの連番を割り当てます。
func Compress(m map[int]int) map[int]int {
	keys := []int{}
	for v := range m {
		keys = append(keys, v)
	}
	sort.Ints(keys)

	for i, v := range keys {
		m[v] = i
	}

	return m
}
