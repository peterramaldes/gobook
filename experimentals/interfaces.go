package main

import (
	"fmt"
	"sort"
)

type Foo interface {
	Run()
}

type Bar struct {
	id    uint
	words []uint
}

func (b *Bar) Run() {
	panic("not implemented") // TODO: Implement
}

func main() {
	a := Bar{id: 1}
	b := Bar{id: 1}
	fmt.Printf("a == b = %+v\n", a == b)
	sort.Sort()
}
