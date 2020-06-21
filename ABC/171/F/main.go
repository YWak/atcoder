package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	K := nextInt()
	L := len(nextString())

	initFact(K + L)
	n := mint(0)
	p1 := mint(1)
	p2 := mint(25).pow(mint(K))
	// iはS[N]以降の文字数
	for i := 0; i <= K; i++ {
		a := L - 1 + K - i
		b := L - 1
		c := facts[a].div(facts[b]).div(facts[a-b])
		m := p1.mul(p2).mul(c)
		n = n.add(m)
		p1 = p1.mul(26)
		p2 = p2.div(25)
	}

	fmt.Println(n)
}

var facts []mint // facts[i] = i!

//var ifacts []mint // ifacts[i] = i!の逆元

// nの階乗までを初期化します
func initFact(n int) {
	facts = make([]mint, n+1)
	facts[0] = mint(1)

	for i := mint(1); i <= mint(n); i++ {
		facts[i] = facts[i-1].mul(i)
	}
}

type mint int64

var mod = mint(1e9 + 7)

// add は a + bを返します
func (a mint) add(b mint) mint {
	return (a + b) % mod
}

// sub は a - bを返します
func (a mint) sub(b mint) mint {
	return (a - b + mod) % mod
}

// mul は a * bを返します
func (a mint) mul(b mint) mint {
	return (a * (b % mod)) % mod
}

// div は a/bを返します
func (a mint) div(b mint) mint {
	return a.mul(b.inv())
}

// inv は aの逆元を返します
func (a mint) inv() mint {
	// 拡張ユークリッドの互除法
	b := mod
	u := mint(1)
	v := mint(0)
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	return (u + mod) % mod
}

// pow は a ^ bを返します
func (a mint) pow(b mint) mint {
	ans := mint(1)

	for b != 0 {
		if b&1 == 1 {
			ans = ans.mul(a)
		}
		a = a.mul(a)
		b = b >> 1
	}
	return ans
}

var stdin = initStdin()

func initStdin() *bufio.Scanner {
	bufsize := 1 * 1024 * 1024 // 1 MB
	var stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufsize), bufsize)
	stdin.Split(bufio.ScanWords)
	return stdin
}

func nextString() string {
	stdin.Scan()
	return stdin.Text()
}

func nextBytes() []byte {
	stdin.Scan()
	return stdin.Bytes()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextString(), 10, 64)
	return i
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args)
}
