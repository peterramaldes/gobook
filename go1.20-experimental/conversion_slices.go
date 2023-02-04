package main

import "fmt"

func main() {
	s := make([]byte, 2, 4)
	fmt.Printf("%v\n", s)
	a0 := [0]byte(s)
	fmt.Printf("%v\n", a0)
}
