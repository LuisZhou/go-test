package main

import (
	"fmt"
	"math/rand"
	//"time"
)

type Worker struct {
	id int
}

func main() {
	// A channel is a communication pipe between goroutines which is used to pass data. In other words,
	// a goroutine that has data can pass it to another goroutine via a channel. The result is that, at any point in time,
	// __only one goroutine__ (no matter writer or reader) has access to the data.

	// The final thing to know before we look at our first example is that receiving and sending to and from a channel is
	// blocking. That is, when we receive from a channel, execution of the goroutine won't continue until data is
	// available. Similarly, __when we send to a channel, execution won't continue until the data is received__
	// (only one buffer).

	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	i := 1

	// Only one buffer. Only when the old data is received, the sender can send new data.
	// you can make the sender is more quickly than the receive to see the result.
	for {
		fmt.Printf("send %d \n", i)
		c <- rand.Int()
		//time.Sleep(time.Millisecond * 50)
		i++
	}
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		// time.Sleep(time.Millisecond * 500)
	}
}
