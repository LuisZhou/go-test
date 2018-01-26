package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	// type IP []byte : IP type is just a byte array.
	// but golang bind some operation to it, such as:
	// func (ip IP) DefaultMask() IPMask
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
		// example: 127.0.0.1, we think it may be '7F 00 00 01', but it print:
		// 00 00 00 00 00 00 00 00 00 00 FF FF 7F 00 00 01
		// means this is ipv6	compatible.
		for _, c := range addr {
			fmt.Printf("%02X ", c)
		}
		fmt.Printf("\n")
		// type IPMask []byte
		// this output just: 'ff000000'
		// var mask net.IPMask = net.IPv4Mask(0xff, 0, 0, 0)
		// fmt.Println("test mask is ", mask.String())
	}
	os.Exit(0)
}
