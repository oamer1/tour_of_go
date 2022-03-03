package main

import (
	"fmt"
)

// Newton's method
func Sqrt(x float64) float64 {
	// starting guess
	z := 1.0
	for i := 0;i < 10;i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		
	}
	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
