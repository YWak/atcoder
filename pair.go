//lint:file-ignore U1000 using template
package main

import (
	"sort"
)

// ==================================================
// 構造
// ==================================================
type pair struct {
	a, b int
}

// sortPairsはpairの配列をソートします。
func sortPairs(array *[]*pair) {
	sort.Slice(*array, func(i, j int) bool {
		pi, pj := (*array)[i], (*array)[j]
		if pi.a != pj.a {
			return pi.a < pj.a
		}
		return pi.b < pj.b
	})
}
