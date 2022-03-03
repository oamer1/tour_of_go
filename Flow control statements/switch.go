package main

import (
	"fmt"
	"runtime"
)

// A switch statement is a shorter way to write a sequence of if - else statements
// Similar to C, Java but it it only runs selected case not what follows
// Provide break for you
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
