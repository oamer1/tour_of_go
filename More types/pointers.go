package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
// & operator generates a pointer to its operand
// The * operator denotes the pointer's underlying value.
// Unlike C no pointers arithmetic
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	// This is known as "dereferencing" or "indirecting
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
