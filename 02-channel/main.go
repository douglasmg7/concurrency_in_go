package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("Sending:", i)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("Received: %d\n", v)
		// fmt.Println("Received:", v)
	}
}
