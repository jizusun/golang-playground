package main

func main() {
	c := make(chan struct{})
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		c <- struct{}{}
	}()
	final := <-c
	println(final)
}
