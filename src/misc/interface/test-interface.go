package main

import (
	"fmt"
	"reflect"
)

type CallBack func(ret interface{}, err error)

func main() {
	var i1 interface{} = call0()
	var i2 interface{} = callN()
	fmt.Println(i1, i2)

	callV()

	var cb CallBack = func(ret interface{}, err error) {
		fmt.Println("callback")
	}
	cb(0, nil)

	// var cb2 interface{} = cb1
	// switch cb1.(type) {
	// case func(interface{}, error):
	// default:
	// 	panic("definition of callback function is invalid")
	// }

	var cb1 CallBack = nil
	_ = cb1

}

// test nil can cost to interface{}
func call0() interface{} {
	return nil
}

func callN() interface{} {
	var i []interface{} = []interface{}{1, 2, 3}
	return i
}

func callV(args ...interface{}) {
	fmt.Printf("the type of r is %s %T\n", reflect.TypeOf(args), args)
}

func cb1(ret interface{}, err error) {

}
