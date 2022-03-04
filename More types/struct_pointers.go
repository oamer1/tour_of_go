package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	p := &v
	// without the explicit dereference.
	// (*p).X
	p.X = 1e9
	fmt.Println(v)
}
