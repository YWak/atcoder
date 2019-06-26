package utils

type pair struct {
	left  int64
	right int64
}
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Less(i, j int) bool {
	return p[i].left < p[j].left
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
