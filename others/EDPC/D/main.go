package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    N := nextInt()
    W := nextInt64()
    weights := make([]int64, N+1)
    values := make([]int64, N+1)

    dp := make([][]int64, N+1)

    for i := 0; i < N; i++ {
        weights[i] = nextInt64()
        values[i] = nextInt64()

        dp[i] = make([]int64, W+1)
    }
    dp[N] = make([]int64, W+1)
    for i := 0; i < N; i++ {
        for w := int64(0); w <= W; w++ {
            if w >= weights[i] {
                dp[i+1][w] = max(dp[i][w-weights[i]] + values[i], dp[i][w])
            } else {
                dp[i+1][w] = dp[i][w]
            }
        }
    }
    fmt.Println(dp[N][W])
}

func max(a, b int64) int64 {
    if a > b {
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
