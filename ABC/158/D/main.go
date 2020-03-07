package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S1 := nextBytes()
	Q := nextInt()

	// 逆順の文字列を準備する
	N := len(S1)
	S2 := make([]byte, N)
	for i := 0; i < N; i++ {
		S2[i] = S1[N-1-i]
	}

	heads := make([]byte, 0, Q)
	tails := make([]byte, 0, Q)

	for i := 0; i < Q; i++ {
		T := nextInt()

		if T == 1 {
			// 反転
			heads, tails = tails, heads
			S1, S2 = S2, S1
		} else {
			// 文字追加
			F := nextInt()
			C := nextBytes()

			if F == 1 {
				heads = append(heads, C[0])
			} else {
				tails = append(tails, C[0])
			}
		}
	}

	for i := len(heads) - 1; i >= 0; i-- {
		// headsは逆順に表示
		fmt.Printf("%c", heads[i])
	}
	for i := 0; i < len(S1); i++ {
		fmt.Printf("%c", S1[i])
	}
	for i := 0; i < len(tails); i++ {
		// tailsは正順に表示
		fmt.Printf("%c", tails[i])
	}
	fmt.Println()
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
