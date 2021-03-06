// basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(basename2(os.Args[1]))
	} else {
		fmt.Println("Usage: basename <filename>")
	}
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
