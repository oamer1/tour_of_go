/*
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
*/
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	// for methods as a convenience,
	// (&v).Scale(5) same as v.Scale(5) since Scale has pointer receiver

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
