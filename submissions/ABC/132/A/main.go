package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := nextString()
	c1 := 0
	c2 := 0
	found := false
	var second byte

	for i := 0; i < 4; i++ {
		if S[i] == S[0] {
			c1++
		} else {
			if !found {
				second = S[i]
				found = true
			}
			if S[i] == second {
				c2++
			}
		}
	}

	if c1 == 2 && c2 == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
