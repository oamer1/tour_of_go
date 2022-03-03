// Every Go program is made up of packages.
package main

import (
	"fmt"
	"math/rand" // package name is rand
)

// Package entrypoint
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
