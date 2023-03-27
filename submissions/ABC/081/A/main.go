package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	stdin.Scan()
	line1 := stdin.Text()

	c := 0

	for i := 0; i < len(line1); i++ {
		if line1[i] == '1' {
			c++
		}
	}

	fmt.Printf("%d\n", c)
}
