package main

import (
	"fmt"
	"reflect"
)

func main() {
	// At the basic level, reflection is just a mechanism to examine the type and value pair stored inside
	// an interface variable.
	var x float64 = 3.4

	// no matter the interface or other type pass to reflect.TypeOf, every thing just ok.
	// internal implement:

	// TypeOf returns the reflection Type of the value in the interface{}.
	// func TypeOf(i interface{}) Type

	// When we call reflect.TypeOf(x), x is first stored in an empty interface, which is then passed as the argument;
	// reflect.TypeOf unpacks that empty interface to recover the type information.
	// reflect.TypeOf return reflect.Type
	fmt.Println("type:", reflect.TypeOf(x))

	// (We call the String method explicitly because by default the fmt package digs into a reflect.Value to show the
	// concrete value inside. The String method does not.)

	// reflect.ValueOf return reflect.Value

	// value: <float64 Value>
	fmt.Println("value:", reflect.ValueOf(x).String())
	// 3.4
	fmt.Println("value:", reflect.ValueOf(x))

	// we know x pass to interface, and ValueOf return the internal value of the interface.
	// value is same as x. but not the x itself.

	// method of reflect.Value
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// how golang define unaddressable value?
	// panic: reflect: reflect.Value.SetFloat using unaddressable value
	// v.SetFloat(123)

	// First, to keep the API simple, the "getter" and "setter" methods of Value operate on the largest type that can hold
	// the value: int64 for all the signed integers, for instance. That is, the Int method of Value returns an int64 and
	// the SetInt value takes an int64; it may be necessary to convert to the actual type involved
	var y uint8 = 'x'
	v1 := reflect.ValueOf(y)
	fmt.Println("type:", v1.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v1.Kind() == reflect.Uint8) // true.
	y = uint8(v1.Uint())                                       // v.Uint returns a uint64.

	// Second, careful of the Kind()
	// The second property is that the Kind of a reflection object describes the underlying type, not the static type.
	// If a reflection object contains a value of a user-defined integer type, as in
	type MyInt int
	var x2 MyInt = 7
	v2 := reflect.ValueOf(x2)                                      //x2 is MyInt
	fmt.Println("type:", v2.Kind(), reflect.TypeOf(v2), v2.Type()) //type: int reflect.Value main.MyInt
}
