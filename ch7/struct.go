package main

import "fmt"

type IntSet struct{}

func (*IntSet) String() string {
	return "call String()"
}

func main() {
	// IntSet{}.String() // compile error

	var s IntSet
	s.String()

	var _ fmt.Stringer = &s
	// var _ fmt.Stringer = s // compile error
}
