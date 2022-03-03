package main

import (
	"fmt"
	"math"
)

// recursive function
func sqrt(x float64) string {
	// Similar to for loop no parenthesis needed
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
