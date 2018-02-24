package main

import (
	"fmt"
)

// reference:
// https://stackoverflow.com/questions/24096026/in-go-what-happens-if-you-write-to-closed-channel-can-i-treat-channels-as-deter

// If you write to a closed channel, your program will panic.
func main() {
	ch := make(chan int)
	close(ch)
	r, ok := <-ch      // read is just all right.
	fmt.Println(r, ok) // 0 false, no block.
	// ch <- 1 // this will panic

	ch2 := make(chan []byte)
	close(ch2)
	r2, ok2 := <-ch2     // read is just all right.
	fmt.Println(r2, ok2) // 0 false, no block.
}
