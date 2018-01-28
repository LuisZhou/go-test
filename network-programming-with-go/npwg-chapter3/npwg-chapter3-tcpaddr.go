package main

import (
	"fmt"
	"net"
	"os"
)

// The  type TCPAddr  is a structure containing an IP and a port
// type TCPAddr struct {
// 	IP   IP
// 	Port int
// }

// example:
// ../../bin/npwg-chapter3-tcpaddr tcp www.baidu.com:80
// 14.215.177.39:80
// Detail: network tcp, ip [0 0 0 0 0 0 0 0 0 0 255 255 14 215 177 39], port 80, zone

// ./npwg-chapter3-tcpaddr tcp [::1]:23
// [::1]:23
// Detail: network tcp, ip [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1], port 23, zone

// Another special case is often used for servers, where the host address is zero, so that the TCP address is really
// just the port name, as in  ":80"  for an HTTP server.
// ./npwg-chapter3-tcpaddr tcp :23
// :23
// Detail: network tcp, ip [], port 23, zone

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type ip/hostname with port\n", os.Args[0])
		os.Exit(1)
	}
	net_type := os.Args[1]
	service_with_port := os.Args[2]

	// net  is one of  "tcp" ,  "tcp4"  or  "tcp6"
	// addr  is a string composed of a host name or IP address, followed by the port number after a  ":"
	// the address is an IPv6 address, which already has colons in it, then the host part must be enclosed in square brackets, such as  "[::1]:23"
	tcp_addr, _ := net.ResolveTCPAddr(net_type, service_with_port)
	if tcp_addr == nil {
		fmt.Println("Invalid type or hostname")
		os.Exit(1)
	}
	fmt.Println(tcp_addr.String()) // to human readable.
	fmt.Printf("Detail: network %s, ip %d, port %d, zone %s\n", tcp_addr.Network(), tcp_addr.IP, tcp_addr.Port, tcp_addr.Zone)
	os.Exit(0)
}
