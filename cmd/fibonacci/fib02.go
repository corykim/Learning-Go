// uses a standard loop
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var err error

	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "[n]]\n")
		os.Exit(1)
	}

	if n, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result uint64 = 1
	var prev1, prev2 uint64 = 1, 1

	for i := 2; i <= n; i++ {
		if i < 3 {
			result += 1
		} else {
			result = prev1 + prev2
			prev2 = prev1
			prev1 = result
		}
	}
	fmt.Println(result)
}
