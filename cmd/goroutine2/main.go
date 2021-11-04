package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int, 1)
	ch <- 888

	go func() {
		val := <-ch
		fmt.Println(val)
		wg.Done()
	}()
	wg.Wait()

}
