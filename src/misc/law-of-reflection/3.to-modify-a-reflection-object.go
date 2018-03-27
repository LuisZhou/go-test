package main

import (
	"fmt"
	"reflect"
)

type Hello struct {
}

func main() {
	// reflect of interface
	x := 1.0
	var i interface{} = x
	v := reflect.ValueOf(i)
	fmt.Printf("value of interface is %s\n", v) // that is ok.

	v = reflect.ValueOf(x)
	// we pass a value not a ptr.
	// v.SetFloat(7.1) // Error: will panic.
	// if the statement were allowed to succeed, it would not update x, even though v looks like it was created from x.
	// Instead, it would update the copy of x stored inside the reflection value and x itself would be unaffected. That
	// would be confusing and useless, so it is illegal, and settability is the property used to avoid this issue.

	// Settability is a bit like addressability, but stricter. It's the property that a reflection object can modify the
	// actual storage that was used to create the reflection object. Settability is determined by whether the reflection
	// object holds the original item.

	// we pass a copy of x to reflect.ValueOf, so the interface value created as the argument to reflect.ValueOf is a copy
	// of x, not x itself.
	fmt.Println("settability of v:", v.CanSet())

	var x1 float64 = 3.4
	p := reflect.ValueOf(&x1) // Note: take the address of x.
	// type of p: *float64
	fmt.Println("type of p:", p.Type())
	// The reflection object p isn't settable, but it's not p we want to set, it's (in effect) *p.
	fmt.Println("settability of p:", p.CanSet())
	// p.SetFloat(7.1)
	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x1)

	// Just keep in mind that reflection Values need the address of something in order to modify what they represent.

	// type Hello is not an expression
	// v = reflect.ValueOf(Hello)
	// var ty interface{} = Hello
	// fmt.Println(ty.Name())

	// Array, Chan, Map, Ptr, or Slice

	// only interface type can do type assert.
	// non-interface type reflect.Value on left
	// msg := reflect.New(i).Interface()
	// proto.Unmarshal(data, msg.(proto.Message))

	// carefule, Value and Type has methed: Elem()
}
