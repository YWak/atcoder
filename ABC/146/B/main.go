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

	alpha := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	dict := map[byte]byte{}

	for i := 0; i < len(alpha); i++ {
		dict[alpha[i]] = alpha[(i+N)%26]
	}
	for i := 0; i < len(S); i++ {
		fmt.Printf("%c", dict[S[i]])
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
