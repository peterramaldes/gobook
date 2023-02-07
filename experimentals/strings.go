package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "hello, world"
	fmt.Println(s[0], string(s[7]))
	fmt.Println(s[0:5])

	a := "alface"
	b := "Zorro"
	c := "zebra"
	d := "ábaco"

	vet := []string{a, b, c, d}

	sort.Strings(vet)
	fmt.Println(vet) // [Zorro alface zebra ábaco]

	fmt.Println("\uFFFD")

	fmt.Println(`Go is a tool for managing Go source code.

	Usage:
		go command [args]
	...`)
}
