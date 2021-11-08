package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	go func() {
		fmt.Println("inner ping")
	}()

	messages <- "ping"
	messages <- "pong"
	msg := <-messages
	fmt.Println(msg)
	//fmt.Println(msg)
	messages <- "ping"
	messages <- "pong"
	// msg := <-messages
	// println(msg)
	// msg = <-messages
	// println(msg)
}
