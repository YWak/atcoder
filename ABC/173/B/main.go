package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	ac := 0
	wa := 0
	tle := 0
	re := 0

	for i := 0; i < N; i++ {
		s := nextString()

		if s == "AC" {
			ac++
		}
		if s == "WA" {
			wa++
		}
		if s == "TLE" {
			tle++
		}
		if s == "RE" {
			re++
		}
	}
	fmt.Printf("AC x %d\nWA x %d\nTLE x %d\nRE x %d\n", ac, wa, tle, re)
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
