package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding
	// function returns.
	a()
	b()
	c()

	// When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then
	// F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack
	// until all functions in the current goroutine have returned, at which point the program crashes.
	func() {
		d()
		// can not reach if no recover
		fmt.Println("wrapper func continue after called function panic")
	}()
	// can not reach if no recover
	fmt.Println("main function continue after called function panic")
}

// Defer statements allow us to think about closing each file right after opening it, guaranteeing that, regardless of
// the number of return statements in the function, the files will be closed.
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer fmt.Println(i) // print 0
	i++
	return
}

// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// 		Like stack fashion.
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i) // print 3210
	}
}

// 3. Deferred functions may read and assign to the returning function's named return values.
// 		In this example, i is the named return value.
func c() (i int) {
	defer func() { i++; fmt.Println(i) }()
	return 1
}

func d() {
	// will call defer
	defer func() {
		fmt.Println("call defer after panic")
		recover()
	}()
	panic("test defer")
	// can not reach if no recover
	fmt.Println("Statement after panic()")
}
