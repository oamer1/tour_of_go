/* Channels can be buffered. Provide the buffer length as the second
argument to make to initialize a buffered channel:
Sends to a buffered channel block only when the buffer is full.
 Receives block when the buffer is empty.

if buffer is full sender blocks until receiver consumes
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 2 --> deadlock

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
