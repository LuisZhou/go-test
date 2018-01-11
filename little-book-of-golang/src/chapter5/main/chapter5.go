package main

import (
	"fmt"
	"os"
)

type Add func(a int, b int) int

// error interface.
// type error interface {
//   Error() string
// }

func main() {

	// __Error Handling__
	// Go's preferred way to deal with errors is through return values, not exceptions.

	// return errors.New("Invalid count")
	// var EOF = errors.New("EOF")

	// __Defer__
	file, err := os.Open("./chapter5.go") // relative path: base on the position you run the go.
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// __go fmt__
	// When you're inside a project, you can apply the formatting rule to it and all sub-projects
	// go fmt ./...

	// __Initialized If__
	if err := processBeforeIf(); err != nil {
		return
	}

	// __Empty Interface and Conversions__
	add(1, 'a')

	// __Strings and Byte Arrays__
	// Do note that when you use []byte(X) or string(X), you're creating a copy of the data.
	// This is necessary because strings are immutable.
	stra := "the spice must flow"
	byts := []byte(stra) // string to byte slice.
	strb := string(byts) // byte slice to string. more comvert like: int64(xxx)
	fmt.Println(byts)    // this print nuber for the byte.
	fmt.Println(strb)
	fmt.Println(len("椒")) // 3

	testUnicodeString := "椒椒椒"
	//0 26898
	//3 26898
	//6 26898
	for key, value := range testUnicodeString {
		fmt.Println(key, value)
	}

	// key [0, 8]
	for key, value := range []byte(testUnicodeString) {
		fmt.Println(key, value)
	}

	// __Function Type__
	// Functions are first-class types
	// Using functions like this can help decouple code from specific implementations much like we achieve with interfaces
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

func process(adder Add) int {
	return adder(1, 2)
}

func processBeforeIf() error { // careful the return value.
	return nil
}

func add(a interface{}, b interface{}) interface{} {
	// use of .(type) outside type switch
	// fmt.Println(a.(type))
	// fmt.Println(b.(type))

	switch a.(type) {
	case int:
		fmt.Printf("a is now an int and equals %d\n", a.(int))
	case bool, string:
		fmt.Printf("a is now an string/bool and equals %d\n", a)
	default:
		// ...
	}

	return 0
}
