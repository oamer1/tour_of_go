package main

import "fmt"

// return value can be named, variable that are treated at top
// Naked return = return named return values
// Use in short function so readability is preserved
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
