package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	H := nextInt()
	W := nextInt()
	K := nextInt()

	c := make([][]byte, H)

	for i := 0; i < H; i++ {
		c[i] = nextBytes()
	}

	ans := 0

	for i := 0; i < pow2(H); i++ {
		for j := 0; j < pow2(W); j++ {
			k := 0
			for h := 0; h < H; h++ {
				for w := 0; w < W; w++ {
					if nthbit(i, h) == 0 && nthbit(j, w) == 0 && c[h][w] == '#' {
						k++
					}
				}
			}
			if k == K {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func pow2(n int) int {
	return int(math.Pow(float64(2), float64(n)))
}

func nthbit(a int, n int) int { return int((a >> uint(n)) & 1) }

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
