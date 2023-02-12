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
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

func (d database) list(w http.ResponseWriter, _ *http.Request) {
	for i, p := range d {
		fmt.Fprintf(w, "%s: %s\n", i, p)
	}
}

func (d database) price(w http.ResponseWriter, req *http.Request) {
	i := req.URL.Query().Get("item")
	price, ok := d[i]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", i)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
