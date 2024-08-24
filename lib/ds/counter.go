package ds

// Counterは整数の多重集合を扱う簡易な実装です。
type Counter map[int]int

func NewCounter() Counter {
	return Counter{}
}

// Addはvalueの個数を増やします。
func (c *Counter) Add(value int) {
	(*c)[value]++
}

// Removeはvalueの個数を減らします。
func (c *Counter) Remove(value int) {
	v := (*c)[value]
	v--
	if v == 0 {
		delete(*c, value)
	} else {
		(*c)[value] = v
	}
}

// Countはvalueの個数を返します。
func (c *Counter) Count(value int) int {
	return (*c)[value]
}

// CountTypesはこの集合に含まれる数値の種類を返します。
func (c *Counter) CountTypes() int {
	return len(*c)
}
