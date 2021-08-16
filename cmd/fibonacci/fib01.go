// uses recursion to calculate fibonacci sequence
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
		fmt.Println(fmt.Sprintln("Usage: ", os.Args[0], "[n]"))
		os.Exit(1)
	}

	if n, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(fib(n))
}

func fib(n int) int {
	if n < 3 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
