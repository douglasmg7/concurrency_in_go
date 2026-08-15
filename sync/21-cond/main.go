package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	var wg sync.WaitGroup

	wg.Go(func() {
		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
			time.Sleep(1 * time.Millisecond)
		}
		c.L.Unlock()

		fmt.Println(sharedRsc["rsc1"])
	})

	// writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}
