// Dup2 show the text for the entire line that appear more than once in the
// stdin, preceded by the count. This program read by the stdin and list of
// named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := fmt.Sprintf("%s\t%s", input.Text(), f.Name())
		counts[key]++
	}
	// NOTE: ignored error that can be occurrs by input.Err()
}
