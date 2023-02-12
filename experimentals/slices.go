package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	fmt.Printf("a == b = %+v\n", a == b) // compile error, slices cannot be comparable
}
