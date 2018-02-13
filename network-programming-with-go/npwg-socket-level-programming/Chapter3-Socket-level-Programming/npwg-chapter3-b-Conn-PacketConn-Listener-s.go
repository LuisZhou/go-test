/* ThreadedIPEchoServer
 */
package main

import (
	"fmt"
	"net"
	"os"
)

// telnet localhost 1201

func main() {

	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// If you want to write a UDP server, then there is an interface  PacketConn  and a method to return an implementation of this:
// func ListenPacket(net, laddr string) (c PacketConn, err os.Error)

// This interface has primary methods  ReadFrom  and  WriteTo  to handle packet reads and writes.
// The Go net package recommends using these interface types rather than the concrete ones. But by using them, you lose
// specific methods such as  SetKeepAlive  or  TCPConn  and  SetReadBuffer  of  UDPConn , unless you do a type cast.
// It is your choice.
