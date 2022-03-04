package main

import "fmt"

// An array has a fixed size. A slice, on the other hand, is a dynamically-sized
// flexible view into the elements of an array. In practice, slices are much more common than arrays.
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice
	// The type []T is a slice with elements of type T
	// Select element at index 1 up to not including index 4
	var s []int = primes[1:4]
	fmt.Println(s)
}
