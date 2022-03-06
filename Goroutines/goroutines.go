package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// Goroutines run in the same address space, so access to shared memory must be synchronized.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("World")
	say("Hello")
}
