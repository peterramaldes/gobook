package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i, p := range d {
		fmt.Fprintf(w, "%s: %s\n", i, p)
	}
}
