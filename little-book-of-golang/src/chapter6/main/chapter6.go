// Concurrency
// goroutines and channels

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//var counter = 1

var (
	counter = 0
	lock    sync.Mutex
	rwlock  sync.RWMutex
)

func main() {
	// __Goroutines__
	// A goroutine is similar to a thread, but it is scheduled by Go, not the OS.
	// Code that runs in a goroutine can run concurrently with other code

	// property:
	// 1. easy to create and have little overhead
	// 2. M:N threading model (a goroutine has a fraction of overhead (a few KB) than OS threads.
	//		On modern hardware, it's possible to have millions of goroutines.)
	// 3. we need to coordinate our code, do not use sleep.

	fmt.Println("start")
	go process()
	// goroutines with anonymous function.
	go func() {
		fmt.Println("processing in anonymous function")
	}()
	// That's because the main process exits before the goroutine gets a chance to execute
	// (the process doesn't wait until all goroutines are finished before exiting)
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")

	// __Synchronization__
	// channel use to coordinated goroutines.
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)

	for i := 0; i < 10; i++ {
		go incrReader(i)
	}
	for i := 0; i < 10; i++ {
		go incrWrite(i)
	}
	time.Sleep(time.Millisecond * 10)

	// Furthermore, part of concurrent programming isn't so much about serializing access across the narrowest possible
	// piece of code; it's also about coordinating multiple goroutines.

	// __Channels__
	// set channel.go
}

func process() {
	// fmt.Println("processing")
	// this will exit when the main exist.
	for {
		fmt.Println("processing")
		time.Sleep(time.Millisecond * 10)
	}
}

func incr() {
	// one goroutine would be reading counter while another writes to it.
	// a simple line of code, but it actually gets broken down into multiple assembly statements

	// The only concurrent thing you can safely do to a variable is to read from it.
	// You can have as many readers as you want, but writes need to be synchronized.
	lock.Lock()
	defer lock.Unlock()

	counter++
	fmt.Println(counter)
}

func incrReader(i int) {
	rwlock.RLock()
	defer rwlock.RUnlock()
	fmt.Println("I'm reader " + strconv.Itoa(i))
}

func incrWrite(i int) {
	rwlock.Lock()
	defer rwlock.Unlock()
	counter++
	fmt.Println("I'm writer " + strconv.Itoa(i) + " " + strconv.Itoa(counter))
}
