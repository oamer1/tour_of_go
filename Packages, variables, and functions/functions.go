package main

import "fmt"

// Type comes after variable name
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
