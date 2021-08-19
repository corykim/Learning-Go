package main

import (
	"fmt"
	"gopl.io/chapter-02/popcount"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("popcount of %v = %v\n", x, popcount.PopCount(x))
	}
}
