package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(basename(os.Args[1:2][0]))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	return s
}
