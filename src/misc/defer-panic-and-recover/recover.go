package main

import "fmt"

// source: https://blog.golang.org/defer-panic-and-recover

// For a real-world example of panic and recover, see the json package from the Go standard library.
// http://golang.org/pkg/encoding/json/

// The convention in the Go libraries is that even when a package uses panic internally, its external API still presents
// explicit error return values.

func main() {
	i := f()
	fmt.Printf("Returned normally from f ret: %d.\n", i)
}

func f() (i int) {
	// If we remove the deferred function from f the panic is not recovered and reaches the top of the goroutine's call
	// stack, terminating the program. This modified program will output:
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			i = 2 // this make sense.
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
	return 1
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v %v", "Hello world", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// Calling g.
// Printing in g 0
// Printing in g 1
// Printing in g 2
// Printing in g 3
// Panicking!
// Defer in g 3
// Defer in g 2
// Defer in g 1
// Defer in g 0
// Recovered in f Hello world 4
// Returned normally from f.
