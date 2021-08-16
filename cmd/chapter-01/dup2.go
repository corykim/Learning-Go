package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		// read from stdin
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			_ = f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%4d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	// NOTE: ignoring potential errors from input.Err()
	for input.Scan() {
		counts[input.Text()]++
	}
}
