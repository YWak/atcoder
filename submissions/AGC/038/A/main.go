package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    H := nextInt()
    W := nextInt()
    A := nextInt()
    B := nextInt()

    w0 := make([]int, H)
    w1 := make([]int, H)

    mat := make([][]int, H)
    for i := 0; i < H; i++ {
        mat[i] = make([]int, W)
    }
    for w := 0; w < W; w++ {
        h1 := 0

        for h := 0; h < H; h++ {
            if w0[h] == A && w1[h] == A {
                if h1 == B {
                    // fmt.Println("fill 0")
                    mat[h][w] = 0
                    w0[h]++
                } else {
                    // fmt.Println("fill 1")
                    mat[h][w] = 1
                    w1[h]++
                    h1++
                }
            } else if w0[h] == A {
                // fmt.Println("add 1")
                mat[h][w] = 1
                w1[h]++
                h1++
            } else if w1[h] == A {
                // fmt.Println("add 0")
                mat[h][w] = 0
                w0[h]++
            } else if h1 < B {
                // fmt.Println("add2 1")
                mat[h][w] = 1
                w1[h]++
                h1++
            } else {
                // fmt.Println("add2 0")
                mat[h][w] = 0
                w0[h]++
            }
        }
    }
    for i := 0; i < H; i++ {
        for j := 0; j < W; j++ {
            fmt.Print(mat[i][j])
        }
        fmt.Println()
    }
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
