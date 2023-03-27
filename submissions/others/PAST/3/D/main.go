package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	S := make([][]byte, 5)
	for i := 0; i < 5; i++ {
		S[i] = nextBytes()
	}

	for i := 0; i < N; i++ {
		fmt.Print(guess(S, i))
	}

	fmt.Println()
}

func c(b byte) string {
	return string([]byte{b})
}
func guess(s [][]byte, n int) int {
	b := n*4 + 1

	// for i := 0; i < 5; i++ {
	// 	println(c(s[i][b+0]), c(s[i][b+1]), c(s[i][b+2]))
	// }

	if s[1][b+1] == '#' {
		return 1
	}
	if s[0][b+1] == '.' {
		return 4
	}
	if s[2][b+1] == '.' {
		if s[4][b+1] == '#' {
			return 0
		}
		return 7
	}
	if s[1][b+0] == '#' {
		// 5689
		if s[1][b+2] == '#' {
			// 89
			if s[3][b+0] == '#' {
				return 8
			}
			return 9
		}
		if s[3][b+0] == '#' {
			return 6
		}
		return 5
	}
	if s[3][b+2] == '#' {
		return 3
	}
	return 2
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
