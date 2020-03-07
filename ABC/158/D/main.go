package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := nextBytes()
	Q := nextInt()

	heads := make([]byte, 0)
	tails := make([]byte, 0)

	front := true

	for i := 0; i < Q; i++ {
		T := nextInt()

		if T == 1 {
			// 反転
			front = !front
		} else {
			// 文字追加
			F := nextInt()
			C := nextBytes()

			if (front && F == 1) || (!front && F == 2) {
				heads = append(heads, C[0])
			} else {
				tails = append(tails, C[0])
			}
		}
	}

	str := make([]byte, 0)
	if front {
		for i := len(heads) - 1; i >= 0; i-- {
			str = append(str, heads[i])
		}
		for i := 0; i < len(S); i++ {
			str = append(str, S[i])
		}
		for i := 0; i < len(tails); i++ {
			str = append(str, tails[i])
		}
	} else {
		for i := len(tails) - 1; i >= 0; i-- {
			str = append(str, tails[i])
		}
		for i := len(S) - 1; i >= 0; i-- {
			str = append(str, S[i])
		}
		for i := 0; i < len(heads); i++ {
			str = append(str, heads[i])
		}
	}
	fmt.Println(string(str))
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
