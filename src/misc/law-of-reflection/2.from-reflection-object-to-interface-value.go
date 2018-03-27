package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 1.0
	v := reflect.ValueOf(x)
	y := v.Interface().(float64)
	// The arguments to fmt.Println, fmt.Printf and so on are all passed as empty interface values, which are then
	// unpacked by the fmt package internally just as we have been doing in the previous examples.

	// the empty interface value has the concrete value's type information inside and Printf will recover it.
	// 1.000000 1.000000 1.0e+00
	fmt.Printf("interface of value is %f %f %7.1e\n", v.Interface(), y, v.Interface())
}
