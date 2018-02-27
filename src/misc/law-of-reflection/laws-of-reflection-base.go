package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	type MyInt int

	// The variables i and j have distinct static types and, although they have the same underlying type, they cannot be
	// assigned to one another without a conversion
	var i int
	var j MyInt

	fmt.Printf("type of i is %s\n", reflect.TypeOf(i))
	fmt.Printf("type of j is %s\n", reflect.TypeOf(j))

	// cannot use j (type MyInt) as type int in assignment
	// i = j

	// example of interface is:
	// Reader is the interface that wraps the basic Read method.
	// type Reader interface {
	// 	Read(p []byte) (n int, err error)
	// }

	// // Writer is the interface that wraps the basic Write method.
	// type Writer interface {
	// 	Write(p []byte) (n int, err error)
	// }

	// One important category of type is interface types, which represent fixed sets of methods. An interface variable can
	// store any concrete (non-interface) value as long as that value implements the interface's methods.
	var r io.Reader
	r = os.Stdin           //implement io.Reader
	r = bufio.NewReader(r) //implement io.Reader
	r = new(bytes.Buffer)  //implement io.Reader

	// whatever concrete value r may hold, r's type is always io.Reader: Go is statically typed and the static type of r
	// is io.Reader

	// beside the 3th, others all print bytes.Buffer
	fmt.Printf("the type of r is %s %s %s %T\n", reflect.TypeOf(r), typeof(r), reflect.TypeOf(r).String(), r)

	// is also a type.
	// It represents the empty set of methods and is satisfied by any value at all, since any value has zero or more methods.
	// interface{}

	// ## The representation of an interface (not just for interface{}, but all others interface)

	// detailed blog post about the representation of interface values in Go
	// http://research.swtch.com/2009/12/go-data-structures-interfaces.html

	var reader io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("can not open /dev/tty")
		return
	}
	// reader contains, schematically, the (value, type) pair, (tty, *os.File)
	reader = tty

	// type assertion;
	var w io.Writer
	// what it asserts is that the item inside r also implements io.Writer, and so we can assign it to w
	// w contain (tty, *os.File)
	w = reader.(io.Writer)

	// The static type of the interface determines what methods may be invoked with an interface variable, even though the
	// concrete value inside may have a larger set of methods.
	var empty interface{}
	// (We don't need a type assertion here because it's known statically that w satisfies the empty interface. In the
	// example where we moved a value from a Reader to a Writer, we needed to be explicit and use a type assertion because
	//Writer's methods are not a subset of Reader's.)
	empty = w
	_ = empty

	// One important detail is that the pair inside an interface always has the form (value, concrete type) and cannot
	// have the form (value, interface type). Interfaces do not hold interface values.
}

func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case io.Reader:
		return "io.Reader"
	//... etc
	default:
		_ = t
		return "unknown"
	}

	// ref:
	// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
	// https://gobyexample.com
	// https://golang.org/pkg/reflect/
}
