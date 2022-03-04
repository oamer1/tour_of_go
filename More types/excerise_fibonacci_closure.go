package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// fib(n) = fib(n-1) + fib(n-2)
	n_0, n_1 := 0, 1
	return func() (ret int) {
		ret, n_0, n_1 = n_0, n_1, n_0+n_1
		return ret

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
