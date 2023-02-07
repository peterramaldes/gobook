package main

import (
	"fmt"
	"path"
)

func main() {
	foo := path.Dir("../a/b/d/c/a.java")
	fmt.Println(foo)
}
