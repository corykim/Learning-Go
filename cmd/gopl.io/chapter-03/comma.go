// formats an integer with commas as thousands separators
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(comma(os.Args[1]))
	} else {
		fmt.Println("Usage: comma <integer>")
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
