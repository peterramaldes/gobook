package main

import (
	"fmt"
)

func main() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		// dir := d // NOTE: necessary, why?
		// os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			fmt.Printf("d: %#v\n", d)
			// fmt.Printf("dir: %#v\n", dir)
			// os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func tempDirs() []string {
	return []string{
		"/tmp/ch5/func/foo",
		"/tmp/ch5/func/bar",
		"/tmp/ch5/func/baz",
	}
}
