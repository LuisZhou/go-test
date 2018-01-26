package main

import (
	"fmt"
	"net"
	"os"
)

// example:
// ../../bin/npwg-chapter3-tcpaddr tcp www.baidu.com:80
// 14.215.177.39:80
// Detail: network tcp, ip [0 0 0 0 0 0 0 0 0 0 255 255 14 215 177 39], port 80, zone

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type ip/hostname with port\n", os.Args[0])
		os.Exit(1)
	}
	net_type := os.Args[1]
	service_with_port := os.Args[2]

	tcp_addr, _ := net.ResolveTCPAddr(net_type, service_with_port)
	if tcp_addr == nil {
		fmt.Println("Invalid type or hostname")
		os.Exit(1)
	}
	fmt.Println(tcp_addr.String())
	fmt.Printf("Detail: network %s, ip %d, port %d, zone %s\n", tcp_addr.Network(), tcp_addr.IP, tcp_addr.Port, tcp_addr.Zone)
	os.Exit(0)
}
