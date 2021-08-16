// uses recursion to calculate fibonacci sequence
package main

import (
	"fmt"
)

func main() {
	n := 50
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n < 3 {
		return 1
	} else {
		return fib(n - 1) + fib(n - 2)
	}
}
