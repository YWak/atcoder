package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	A := []int{
		nextInt(),
		nextInt(),
		nextInt(),
	}
	sum := A[0] + A[1] + A[2]
	if sum == 0 {
		fmt.Println("No")
		return
	}

	S := make([]int, N)
	ans := make([]int, N)

	for i := 0; i < N; i++ {
		s := nextString()

		if s == "AB" {
			S[i] = 1
		} else if s == "AC" {
			S[i] = 2
		} else {
			S[i] = 3
		}
	}

	for i := 0; i < N; i++ {
		var a, b, c int
		s := S[i]
		if s == 1 {
			a = 0
			b = 1
			c = 2
		} else if s == 2 {
			a = 0
			b = 2
			c = 1
		} else {
			a = 1
			b = 2
			c = 0
		}
		if A[a] == 0 && A[b] == 0 {
			fmt.Println("No")
			return
		}
		if sum == 2 && (A[a] == 1 && A[b] == 1) && (i != N-1 && S[i] != S[i+1]) {
			if S[i+1] == a+c {
				ans[i] = a
				A[a]++
				A[b]--
			} else {
				ans[i] = b
				A[a]--
				A[b]++
			}
		} else if A[a] < A[b] {
			ans[i] = a
			A[a]++
			A[b]--
		} else {
			ans[i] = b
			A[b]++
			A[a]--
		}
		// fmt.Println(S[i], ans[i], A[0], A[1], A[2])
	}

	fmt.Println("Yes")
	a := []string{"A", "B", "C"}

	for i := 0; i < N; i++ {
		fmt.Println(a[ans[i]])
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
