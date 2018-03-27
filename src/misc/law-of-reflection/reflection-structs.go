package main

import (
	"fmt"
	"reflect"
)

// A common way for this situation to arise is when using reflection to modify the fields of a structure. As long as we
// have the address of the structure, we can modify its fields.

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f) // we pass fmt, f.Interface() not f itself.
	}

	// (Why not fmt.Println(v)? Because v is a reflect.Value; we want the concrete value it holds.)
	// but the result is the same.

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
