package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := nextBytes()

	if solve(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func count(s []byte) [10]int {
	n := [10]int{}
	for i := 0; i < len(s); i++ {
		n[s[i]-'0']++
	}

	return n
}

func solve(S []byte) bool {

	if len(S) == 1 {
		return S[0] == '8'
	}
	if len(S) == 2 {
		a := (S[1]-'0')*10 + (S[0] - '0')
		b := (S[0]-'0')*10 + (S[1] - '0')

		return a%8 == 0 || b%8 == 0
	}

	n := count(S)

loop:
	for i := 0; i <= 999; i += 8 {
		s := []byte(fmt.Sprintf("%03d", i))
		n1 := count(s)

		if n1[0] > 0 {
			continue
		}

		for i := 0; i < 10; i++ {
			if n[i] < n1[i] {
				continue loop
			}
		}
		return true
	}

	return false
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
