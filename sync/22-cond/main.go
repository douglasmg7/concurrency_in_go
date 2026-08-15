package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Go(func() {
		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		c.L.Unlock()
		fmt.Println(sharedRsc["rsc1"])
	})

	wg.Go(func() {
		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		c.L.Unlock()
		fmt.Println(sharedRsc["rsc2"])
	})

	// writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()

	wg.Wait()
}
