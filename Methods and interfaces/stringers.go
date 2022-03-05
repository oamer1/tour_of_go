/*
type Stringer interface {
	String() string
}

https://pkg.go.dev/fmt#Stringer
Stringer is implemented by any value that has a String method, which defines the “native” format for that value
*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
