/* GetHeadInfo
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// ../bin/npwg-chapter3-getheadinfo www.github.com:80
// HTTP/1.1 301 Moved Permanently
// Content-length: 0
// Location: https:///
// Connection: close

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port \n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// one of  "tcp4" ,  "tcp6"  or  "tcp"
	// depending on whether you want a TCPv4 connection, a TCPv6 connection or don't care.
	//
	// if you in ipv4 network, and use parameter 'tcp6'
	// Fatal error: dial tcp6 192.30.255.113:80: connect: network is unreachable
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// In this case, we read essentially a single response from the server.
	// This will be terminated by __end-of-file__ on the connection.
	// However, it may consist of several TCP packets, so we need to keep reading till the end of file.

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

// In C, the same behaviour is gained by special values such as  NULL , or  -1 , or zero being returned -
// if that is possible.

// In Java, the same error checking is managed by throwing and catching exceptions,
// which can make the code look very messy

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
