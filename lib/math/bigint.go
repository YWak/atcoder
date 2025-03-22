package math

import (
	"math/big"
)

type BigInt big.Int

func NewBigInt(n int) *BigInt {
	return ptr(big.NewInt(int64(n)))
}

func ptr(v *big.Int) *BigInt {
	p := BigInt(*v)
	return &p
}

func (n *BigInt) raw() *big.Int {
	v := big.Int(*n)
	return &v
}

func (n *BigInt) Add(v *BigInt) *BigInt {
	return ptr(new(big.Int).Add(n.raw(), v.raw()))
}

func (n *BigInt) Addn(v int) *BigInt {
	return n.Add(NewBigInt(v))
}

func (n *BigInt) Sub(v *BigInt) *BigInt {
	return ptr(new(big.Int).Sub(n.raw(), v.raw()))
}

func (n *BigInt) Subn(v int) *BigInt {
	return n.Sub(NewBigInt(v))
}

func (n *BigInt) Mul(v *BigInt) *BigInt {
	return ptr(new(big.Int).Mul(n.raw(), v.raw()))
}

func (n *BigInt) Muln(v int) *BigInt {
	return n.Mul(NewBigInt(v))
}

func (n *BigInt) Div(v *BigInt) *BigInt {
	return ptr(new(big.Int).Div(n.raw(), v.raw()))
}

func (n *BigInt) Divn(v int) *BigInt {
	return n.Div(NewBigInt(v))
}

func (n *BigInt) Mod(v *BigInt) *BigInt {
	return ptr(new(big.Int).Mod(n.raw(), v.raw()))
}

func (n *BigInt) Modn(v int) *BigInt {
	return n.Mod(NewBigInt(v))
}

func (n *BigInt) Pow(v *BigInt) *BigInt {
	return ptr(new(big.Int).Exp(n.raw(), v.raw(), nil))
}

func (n *BigInt) Pown(v int) *BigInt {
	return n.Pow(NewBigInt(v))
}

func (n *BigInt) PowMod(v *BigInt, m *BigInt) *BigInt {
	return ptr(new(big.Int).Exp(n.raw(), v.raw(), m.raw()))
}

func (n *BigInt) GCD(v *BigInt) *BigInt {
	g, _, _ := n.ExGCD(v)
	return g
}

// ExGCD は x*a + y*b = gとなるような最大のgと、そのときのa,bの例を返します。
// x = y = 0のときはg = 0となります。
func (x *BigInt) ExGCD(y *BigInt) (g, a, b *BigInt) {
	var _a, _b big.Int
	_g := new(big.Int).GCD(&_a, &_b, x.raw(), y.raw())
	g, a, b = ptr(_g), ptr(&_a), ptr(&_b)
	return
}

func (n *BigInt) Abs() *BigInt {
	return ptr(new(big.Int).Abs(n.raw()))
}

func (n *BigInt) Sign() int {
	return new(big.Int).Sign()
}

func (n *BigInt) Cmp(v *BigInt) int {
	return n.raw().Cmp(v.raw())
}

func (n *BigInt) Cmpn(v int) int {
	return n.raw().Cmp(big.NewInt(int64(v)))
}

func (n *BigInt) String() string {
	return n.raw().String()
}

func (n *BigInt) Int() int {
	return int(n.raw().Int64())
}
