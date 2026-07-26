package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 6)

	go func() {
		defer close(ch)

		for i := range 6 {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
