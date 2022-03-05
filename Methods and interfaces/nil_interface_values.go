// A nil interface value holds neither value nor concrete type.

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// Runtime error no type inside the interface tuple to indicate which concrete method to call.
	// (<nil>, <nil>)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
