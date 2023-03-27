package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := nextInt()
	a := nextInt()
	b := nextInt()

	k := max(a, b)
	perms := make([]mint, b+1)
	facts := make([]mint, b+1)

	perms[1] = mint(n)
	facts[1] = mint(1)
	for i := 2; i <= k; i++ {
		perms[i] = perms[i-1].mul(mint(n - i + 1))
		facts[i] = facts[i-1].mul(mint(i))
	}

	// 全組み合わせ - aが入るパターン - bが入るパターン
	ans := mint(2).pow(mint(n)).sub(1)
	ans = ans.sub(perms[a].div(facts[a])).sub(perms[b].div(facts[b]))

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
