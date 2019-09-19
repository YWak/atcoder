package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    N := nextInt()
    dp := make([][3]int, N)

    a := nextInt()
    b := nextInt()
    c := nextInt()
    dp[0] = [3]int{a, b, c}

    for i := 1; i < N; i++ {
        a = nextInt()
        b = nextInt()
        c = nextInt()
        dp[i] = [3]int{max(dp[i-1][1], dp[i-1][2])+a, max(dp[i-1][0], dp[i-1][2])+b, max(dp[i-1][0], dp[i-1][1])+c}
    }

    fmt.Println(max(dp[N-1][0], dp[N-1][1], dp[N-1][2]))
}

func max(a int, b ...int) int {
    l := len(b)
    for i := 0; i < l; i++ {
        if a < b[i] {
            a = b[i]
        }
    }

    return a
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
