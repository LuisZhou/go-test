package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id int
}

func main() {
	c := make(chan int)
	for i := 0; i < 1; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	// + The first available channel is chosen.
	// + If multiple channels are available, one is randomly picked.
	// + If no channel is available, the default case is executed.
	// + If there's no default, select blocks.

	i := 1
	for {
		select {
		case c <- i: // block to wait it
		case <-time.After(time.Millisecond * 50): // block to wait it
			fmt.Printf("timed out %d\n", i)
		}
		time.Sleep(time.Millisecond * 50)
		i++
	}
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 100)
	}
}
