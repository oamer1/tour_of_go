package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// mutate
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// checl presence
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
