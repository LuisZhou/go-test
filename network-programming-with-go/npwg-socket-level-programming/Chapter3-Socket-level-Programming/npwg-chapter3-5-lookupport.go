/* LookupPort
 */

package main

import (
	"fmt"
	"net"
	"os"
)

// example:
// npwg-chapter3-lookupport tcp telnet

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	// same servise has both udp and tcp, can use the same port, such as DNS, use tcp/udp 53.

	// port: unsigned integer between 1 and 65,535

	// net.LookupPort interrogate look up /etc/services in unix.
	// The network argument is a string such as  "tcp"  or  "udp" ,
	// while the service is a string such as  "telnet"  or  "domain"
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port ", port)
	os.Exit(0)
}
