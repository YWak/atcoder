package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	H := nextInt()
	W := nextInt()

	S := make([][]byte, H+1)

	for h := 0; h < H; h++ {
		S[h] = nextBytes()
	}
	S[H] = make([]byte, W)
	for w := 0; w < W; w++ {
		S[H][w] = '#'
	}

	c := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if S[h][w] == '#' {
				continue
			}
			if S[h+1][w] == '.' {
				c++
			}
			if w != W-1 && S[h][w+1] == '.' {
				c++
			}
		}
	}

	fmt.Println(c)
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
	fmt.Fprintln(os.Stderr, args...)
}
