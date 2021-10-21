package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s[0] = 111
	s[2] = 222
	s = append(s, 555, 666, 10000)
	// descrease the capacity, default
	// n := s[2:3:3]
	// more safe
	// n := s[2:3:3]
	// most safe
	n := make([]int, 3, 5)
	copy(n, s)
	n[0] = 100
	// allocate new memory for the slice
	// n[0] = 5
	n = append(n, 888)
	// Java: collection, generics
	fmt.Println(s, len(s), cap(s))
	fmt.Println(n, len(n), cap(n))
}
