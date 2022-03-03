package main

import (
	"fmt"
	"math"
)

// Name is exported if it begins with capital letter
// "unexported" names are not accessible from outside the package.
func main() {
	fmt.Println(math.Pi)
}
