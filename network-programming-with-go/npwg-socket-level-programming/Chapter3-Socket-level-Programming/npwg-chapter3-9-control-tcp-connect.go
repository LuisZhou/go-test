/* SimpleEchoServer
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// client: telnet localhost 1201

// func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
// useful for client.

// const (
//         Nanosecond  Duration = 1
//         Microsecond          = 1000 * Nanosecond
//         Millisecond          = 1000 * Microsecond
//         Second               = 1000 * Millisecond
//         Minute               = 60 * Second
//         Hour                 = 60 * Minute
// )

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close() // we're finished
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {

		// there is setReadDeadline, setWriteDeadline
		conn.SetDeadline(time.Now().Add(time.Second * 5))

		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
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
