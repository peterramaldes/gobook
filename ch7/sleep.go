package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", time.Second*1, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
