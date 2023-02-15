package main

import "fmt"

func main() {
	m := make(chan int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("foo ", i)
			m <- i
		}(i)
		go func(i int) {
			fmt.Println("bar ", i)
		}(i)
	}

	for {
		fmt.Println(<-m)
	}
}

// func main() {
// 	const tam = 50
// 	ch := make(chan string)
// 	for i := 1; i <= tam; i++ {
// 		go func(i int) { ch <- fmt.Sprintf("Foo %s", i) }(i)
// 		// go foo(i, ch) // initiate goroutine
// 	}
// 	for i := 1; i <= tam; i++ {
// 		fmt.Println(<-ch) // receive channel ch
// 	}
// }

// func foo(i int, ch chan<- string) {
// 	ch <- fmt.Sprintf("Foo %s", i)
// }
