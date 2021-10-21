package main

import "fmt"

func main() {
	// ch := make(chan string)

	go func() {
		fmt.Println("goroutine start")
		// ch <- "hello"
		fmt.Println("goroutine end")
	}()
	fmt.Println("main start")
	/*
		When commenting the below line <-ch, the output:
			main start
			main end
	*/
	// <-ch
	/*
		When not commenting the line, the output:
			main start
			goroutine start
			goroutine end
			main end

		Why?

		## Refs:
		- https://stackoverflow.com/a/18660709/3074866
	*/
	fmt.Println("main end")
}
