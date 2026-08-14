package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		// for i := 0; i < 3; i++ {
		for range 3 {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}

	}()

	// TODO: if there is no value on channel, do not block.
	for range 2 {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("Nothing to do.")
		}

		// Do some processing..
		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)
	}
}
