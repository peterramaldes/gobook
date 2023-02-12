package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int               { return len(s) }
func (s StringSlice) Less(i int, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i int, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	// NOTE: take care of the accents when we order words in portuguese
	n := []string{"a", "z", "A", "j", "â", "ã"}

	names := StringSlice(n)

	fmt.Printf("Before Sort = %+v\n", names)
	sort.Strings(n)
	fmt.Printf("After Sort = %+v\n", names)
}
