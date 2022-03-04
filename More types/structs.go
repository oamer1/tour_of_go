package main

import "fmt"

// Structs collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	// You can omit X, Y
	// Vertex{1, 2}
	fmt.Println(Vertex{X: 1, Y: 2})
}
