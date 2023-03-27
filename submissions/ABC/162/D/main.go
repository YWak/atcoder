package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	S := nextBytes()

	cr := make([]int, N+1)
	cr[N] = 0
	cg := make([]int, N+1)
	cg[N] = 0
	cb := make([]int, N+1)
	cb[N] = 0

	// i番目以降のRGBそれぞれの数を数えておく
	for i := N - 1; i >= 0; i-- {
		cr[i] = cr[i+1]
		cg[i] = cg[i+1]
		cb[i] = cb[i+1]

		if S[i] == 'R' {
			cr[i]++
		}
		if S[i] == 'G' {
			cg[i]++
		}
		if S[i] == 'B' {
			cb[i]++
		}
	}
	ans := 0

	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			if S[i] == S[j] {
				continue
			}
			k := j + 1    // j < kより
			kj := 2*j - i // j - i == k - jとなる点

			if S[i] != 'R' && S[j] != 'R' {
				ans += cr[k]
				if kj < N && S[kj] == 'R' {
					ans--
				}
			} else if S[i] != 'G' && S[j] != 'G' {
				ans += cg[k]
				if kj < N && S[kj] == 'G' {
					ans--
				}
			} else if S[i] != 'B' && S[j] != 'B' {
				ans += cb[k]
				if kj < N && S[kj] == 'B' {
					ans--
				}
			}
		}
	}
	fmt.Println(ans)
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
