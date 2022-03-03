package main

import "fmt"

// Go has only one looping construct for
func main() {
	sum := 0
	//  init_statement; condition expression; post statement
	// init_statement visible only in for scope
	// unlike C, JAVA no parentheses in for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
