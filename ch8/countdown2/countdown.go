package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	i, _ := os.Stdin.Read(make([]byte, 1))
	fmt.Println(i)

	abort := make(chan struct{})
	go func() {
		// Waiting for the input from the user
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing coutdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
