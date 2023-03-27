package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	facts := make([]int, N+1)
	counts := make([]int, N+1)
	for i := 1; i <= N; i++ {
		facts[i] = i
	}
	for i := 1; i <= N; i++ {
		if facts[i] == 1 {
			continue
		}
		// debug(i, facts)
		// debug(i, counts)
		f := facts[i]
		for j := i; j <= N; j += i {
			// debug(facts[j], "/", f)
			counts[f]++
			facts[j] = facts[j] / f
		}
	}
	// debug(facts)
	// debug(counts)
	ans := mint(1)
	for i := 2; i <= N; i++ {
		if counts[i] == 0 {
			continue
		}
		ans = ans.mul(mint(counts[i] + 1))
	}

	fmt.Println(ans)
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
