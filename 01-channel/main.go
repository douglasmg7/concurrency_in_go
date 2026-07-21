package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(a, b int) {
		ch <- a + b
	}(1, 2)

	// res := <-ch
	fmt.Printf("Received value: %d\n", <-ch)
}
