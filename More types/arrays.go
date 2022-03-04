package main

import "fmt"

func main() {

	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hi"
	a[1] = "There"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}

// An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry
