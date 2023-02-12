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
	switch r.URL.Path {
	case "/list":
		for i, p := range d {
			fmt.Fprintf(w, "%s: %s\n", i, p)
		}
	case "/price":
		i := r.URL.Query().Get("item")
		price, ok := d[i]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", i)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}
}
