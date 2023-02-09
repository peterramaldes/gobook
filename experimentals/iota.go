package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)
	fmt.Println(FlagPointToPoint)
	fmt.Println(FlagMulticast)
}
