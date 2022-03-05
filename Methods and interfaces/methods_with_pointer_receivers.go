/*
Why using pointer receiver ?!

1- Modify values pointer points to
2- Avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// *Vertex, even though the Abs method needn't modify its receiver.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
