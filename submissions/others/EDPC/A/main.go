package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    N := nextInt()
    H := make([]int, N)
    dp := make([]int, N)

    for i := 0; i < N; i++ {
        H[i] = nextInt()
    }

    dp[0] = 0
    dp[1] = abs(H[0] - H[1])
    for i := 2; i < N; i++ {
        c1 := abs(H[i-1] - H[i])
        c2 := abs(H[i-2] - H[i])
        dp[i] = min(c1 + dp[i-1], c2 + dp[i-2])
    }

    fmt.Println(dp[N-1])
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
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
