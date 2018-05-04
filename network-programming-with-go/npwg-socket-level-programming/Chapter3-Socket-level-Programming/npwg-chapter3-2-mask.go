package main

import (
	"fmt"
	"net"
	"os"
)

// ./bin/npwg-chapter3-2-mask 2001:0:53aa:64c:18d0:a16f:48f1:7b28
// ./bin/npwg-chapter3-2-mask 192.168.163.132

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask() // return IP
	network := addr.Mask(mask) // return IP
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		"\nDefault mask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	os.Exit(0)
}
