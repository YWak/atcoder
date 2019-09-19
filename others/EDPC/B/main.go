package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

func main() {

    N := nextInt()
    K := nextInt()

    H := make([]int, N)
    dp := make([]int, N)

    for i := 0; i < N; i++ {
        H[i] = nextInt()
        dp[i] = math.MaxInt32
    }
    dp[0] = 0
    for i := 1; i < N; i++ {
        for k := 1; k <= K; k++ {
            if i-k < 0 {
                continue
            }
            h := abs(H[i] - H[i-k])
            dp[i] = min(dp[i], dp[i-k]+h)
        }
    }
    fmt.Println(dp[N-1])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
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
