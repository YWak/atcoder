package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := nextInt()
	k := nextInt()

	if n <= k {
		fmt.Println(dupcomb(n, n))
		return
	}

	ans := mint(0) // 移動なし

	for i := 0; i <= k; i++ {
		a := comb(n, i).mul(dupcomb(n-i, i)) // 0人の部屋の数 * 残りの人の配置
		ans = ans.add(a)
	}

	fmt.Println(ans)
}

var facts []mint  // facts[i] = i!
var ifacts []mint // ifacts[i] = i!の逆元

// nの階乗までを初期化します
func initFact(n int) {
	if len(facts) == 0 {
		facts = []mint{mint(1)}
		ifacts = []mint{mint(1)}
	}

	m := mint(n)
	for i := mint(len(facts)); i <= m; i++ {
		facts = append(facts, facts[i-1].mul(i))
		ifacts = append(ifacts, ifacts[i-1].mul(i.pow(mod-2)))
	}
}

// perm はnPkを返します
func perm(n, k int) mint {
	if n < k {
		return 0
	}
	initFact(n)
	return ifacts[n-k].mul(facts[n])
}

// comb はnCkを返します
func comb(n, k int) mint {
	if n == 0 && k == 0 {
		return mint(1)
	} else if n < k || n < 0 {
		return mint(0)
	}
	initFact(n)
	return facts[n].mul(ifacts[n-k]).mul(ifacts[k])
}

// dupcomb はnHkを返します
func dupcomb(n, k int) mint {
	return comb(n+k-1, n-1)
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
