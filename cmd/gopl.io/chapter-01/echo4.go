package main

import (
	"fmt"
	"os"
)

func main() {
	const offset = 1
	for i, arg := range os.Args[offset:] {
		fmt.Println(offset+i, arg)
	}
}
