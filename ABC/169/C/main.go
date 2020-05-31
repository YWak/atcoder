package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := nextBytes()
	B := nextBytes()

	a := make([]int, len(A))
	ans := make([]int, len(A)+10)

	for i := 0; i < len(A); i++ {
		a[len(A)-1-i] = int(A[i] - '0')
	}

	for i := 0; i < len(a); i++ {
		s := a[i] * int(B[3]-'0')
		ans[i+0] += s
	}
	for i := 0; i < len(a); i++ {
		s := a[i] * int(B[2]-'0')
		ans[i+1] += s
	}
	for i := 0; i < len(a); i++ {
		s := a[i] * int(B[0]-'0')
		ans[i+2] += s
	}

	// 繰り上がり
	for i := 0; i < len(ans); i++ {
		if ans[i] >= 10 {
			ans[i+1] += ans[i] / 10
			ans[i] = ans[i] % 10
		}
	}

	first := false

	for i := len(ans) - 1; i >= 2; i-- {
		if ans[i] == 0 && !first {
			continue
		}
		first = true
		fmt.Print(ans[i])
	}
	if !first {
		fmt.Print(0)
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
