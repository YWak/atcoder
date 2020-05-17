package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	N := nextInt()
	a := N % 10

	if a == 3 {
		fmt.Println("bon")
	} else if a == 0 || a == 1 || a == 6 || a == 8 {
		fmt.Println("pon")
	} else {
		fmt.Println("hon")
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
