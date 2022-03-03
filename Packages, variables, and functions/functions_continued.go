package main

import "fmt"

// if variable share a type you can refactor
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
