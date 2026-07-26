package main

import "fmt"

func genMsg(ch chan<- int) {
	defer close(ch)
	// send message on ch1
	for i := range 6 {
		fmt.Println("genMsg: ", i)
		ch <- i
	}
}

func relayMsg(ch1 <-chan int, ch2 chan<- int) {
	defer close(ch2)
	// recv message on ch1
	for v := range ch1 {
		fmt.Println("relayMsg received: ", v)
		// send it on ch2
		ch2 <- v
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	for v := range ch2 {
		fmt.Println("main received: ", v)
	}
}
