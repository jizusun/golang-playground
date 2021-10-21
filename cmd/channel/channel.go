package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	go func() {
		fmt.Println("inner ping")
	}()

	messages <- "ping"
	messages <- "pong"
	_ = <-messages
	fmt.Println(_)
	//fmt.Println(msg)
	messages <- "ping"
	messages <- "pong"
	// msg := <-messages
	// println(msg)
	// msg = <-messages
	// println(msg)
}
