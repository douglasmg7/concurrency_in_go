package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	for i := range 10 {
		wg.Go(func() {

			//TODO: modify so that load function gets called only once.
			once.Do(load)
			fmt.Printf("Go rontine %d\n", i)

		})
	}
	wg.Wait()
}
