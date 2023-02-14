package main

import (
	"log"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)

	for x := range squares {
		log.Println(x)
	}
}

func counter(naturals chan int) {
	log.Println("Start Counter")
	for x := 0; x < 100; x++ {
		naturals <- x
	}
	close(naturals)
	log.Println("Finish Counter")
}

func squarer(naturals chan int, squares chan int) {
	log.Println("Start Squarer")
	for x := range naturals {
		squares <- x * x
	}
	close(squares)
	log.Println("Finish Squarer")
}
