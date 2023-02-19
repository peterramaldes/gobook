package main

import "fmt"

// var lookupTable map[any]string

type genericLookupTable[K comparable, V any] map[K]V

var lookupTable genericLookupTable[any, string]

func main() {
	fmt.Printf("lookupTable = %+v\n", lookupTable)
}
