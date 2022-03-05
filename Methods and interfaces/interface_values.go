package main

/*

Under the hood, interface values can be thought of as a tuple of a value and a concrete type
*/
import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type MyFloat float64

func (f MyFloat) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	// (&{Hello}, *main.T)
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	// (3.141592653589793, main.MyFloat)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
