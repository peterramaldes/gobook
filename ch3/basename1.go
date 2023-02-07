package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(basename(os.Args[1:2][0]))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	return s
}
