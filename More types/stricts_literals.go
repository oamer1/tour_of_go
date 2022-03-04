package main

import "fmt"

type Vertex struct {
	X, Y int
}

// A struct literal denotes a newly allocated struct value by listing the values of its fields.
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{Y: 0}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
