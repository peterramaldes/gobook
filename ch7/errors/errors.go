package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("err = %#v\n", err)
	fmt.Println(os.IsNotExist(err))
}
