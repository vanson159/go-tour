package main

import "fmt"

func main() {
	// array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 6, 8, 9}
	// A slice is a dynamically-size, flexible view into the elements of an array
	// This selects a half-open range which includes the first element, but excludes the last one
	var s []int = primes[3:6]

	fmt.Println(primes)
	fmt.Println(s)
}
