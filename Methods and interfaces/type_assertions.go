package main

import "fmt"

// t := i.(T) type assertion
func main() {
	// Empty inter
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	// 0 false
	fmt.Println(f, ok)

	// Error
	f = i.(float64) // panic
	fmt.Println(f)
}
