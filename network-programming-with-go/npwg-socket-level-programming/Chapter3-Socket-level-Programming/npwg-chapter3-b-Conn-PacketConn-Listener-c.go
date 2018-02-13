package main

// description
//
// The type  Conn  is an interface and both  TCPConn  and  UDPConn  implement this interface. To a large extent you can
// deal with this interface rather than the two types.

/* IPGetHeadInfo
 */

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// Dial

// The net can be any of  "tcp" ,  "tcp4"  (IPv4-only),  "tcp6"  (IPv6-only),  "udp" ,  "udp4"  (IPv4-only),  "udp6"
// (IPv6-only),  "ip" ,  "ip4"  (IPv4-only) and  "ip6"  IPv6-only). It will return an appropriate implementation of the
// Conn  interface.

// ./bin/npwg-chapter3-b-Conn-PacketConn-Listener-c baidu.com:80

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// careful, this no resolveAddr before!
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
