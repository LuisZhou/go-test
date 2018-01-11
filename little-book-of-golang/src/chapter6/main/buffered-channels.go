package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func main() {
	// background: sender is more quickly than the reader.
	// If no worker is available, we want to temporarily store the data in some sort of queue.
	// this just make the writer/sender don't block.
	c := make(chan int, 100)
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
		time.Sleep(time.Millisecond * 50)
		i++
	}
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		// You can see that it grows and grows until it fills up, at which point sending to our channel start to block again.
		fmt.Println(len(c))
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}
