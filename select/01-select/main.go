package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// TODO: mu:tiplex recv on channel - ch1, ch2
	for range 2 {
		select {
		case msg1 := <-ch1:
			fmt.Printf("ch1: %v\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("ch2: %v\n", msg2)
		}
	}

	fmt.Println("Finish.")

}
