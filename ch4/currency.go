package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
	BRL
)

func main() {

	symbol := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RMB: "¥",
		BRL: "R$",
	}

	fmt.Println(RMB, symbol[RMB])
}
