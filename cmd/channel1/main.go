package main

import "fmt"

func main() {
	values := []string{"foo", "bar", "baz"}
	c := make(chan struct{})
	for i, v := range values {
		go func() {
			fmt.Println(i, v)
			fmt.Println("receive")
			<-c
		}()
		fmt.Println("send")
		c <- struct{}{}
	}
}
