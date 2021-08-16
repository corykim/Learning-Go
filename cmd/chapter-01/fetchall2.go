// Fetchall fetches URLs in parallel and reports their times and sizes.
// Also writes the body to a file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch2(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch2(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// write the output to a (sanitized) file
	fnameBase := fmt.Sprintf("fetchall.%s.html", url)
	re := regexp.MustCompile("[^\\w\\.]+")
	filename := re.ReplaceAllString(fnameBase, "_")

	fmt.Println("Filename is", filename)
	outfile, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	dst := io.Writer(outfile)
	nbytes, err := io.Copy(dst, resp.Body)
	_ = outfile.Close()
	_ = resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
