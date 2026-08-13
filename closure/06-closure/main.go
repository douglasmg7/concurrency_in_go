package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		num := i
		wg.Go(func() {
			fmt.Println(num)
		})
	}
	wg.Wait()
}
