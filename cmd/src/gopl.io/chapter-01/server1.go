// Server1 is a minimal http 'echo' server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the http request URL r
func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
