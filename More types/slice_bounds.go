package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	v := s[1:4]
	fmt.Println(v)

	// The default is zero for the low bound and the length of the slice for the high bound.
	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
