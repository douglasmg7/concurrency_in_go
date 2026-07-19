package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for range 3 {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("Bla Bla Bla!")

	// goroutine with anonymous function
	go func() {
		fun("Anonymous.")
	}()

	// goroutine with function value call
	fn := fun
	go fn("Value call")

	// wait for goroutines to end

	fmt.Println("Waiting go routines finish...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("done..")
}
