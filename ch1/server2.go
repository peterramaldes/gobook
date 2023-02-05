// Server2 is simple web server of echo minimal with counter
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)      // each request call handler
	http.HandleFunc("/count", counter) // each request call handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler prints the path of url that are requested
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Println("Entrou em '/'")
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter shows the number of request so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Println("Entrou em '/count'")
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
