package main

import "fmt"

func main() {
	sum := 1
	// Go's while loop !
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}