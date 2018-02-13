/* UDP
 */
package main

// property:
// 1. In a connectionless protocol each message contains information about its origin and destination.
// 2. There is no "session" established using a long-lived socket.
// 3. There is no state maintained by these messages, unless the client or server does so. The messages are not
//		guaranteed to arrive, or may arrive out of order.

// question:
// 1. if datagram limit the size?
// 2. should and how the server/client decide if re-send the datagram?

// satuation:
// The most common situation for a client is to send a message and hope that a reply arrives.
// The most common situation for a server would be to receive a message and then send one or more replies back to that client.
// In a peer-to-peer situation, though, the server may just forward messages to other peers.

// important:
// The major difference between TCP and UDP handling for Go is how to deal with packets arriving from possibly multiple
// clients, without the cushion of a TCP session to manage things.

// The client for a UDP time service doesn't need to make many changes, just changing  ...TCP...  calls to  ...UDP...  calls

// reference:
// https://golang.org/pkg/net

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	// careful:
	// all client share one for-loop
	for {
		handleClient(conn)
	}
}

// careful:
// ReadFromUDP return udp address too.
func handleClient(conn *net.UDPConn) {
	var buf [512]byte

	// if err, you should decide what to do next.
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()

	// if err, you should decide what to do next.
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
