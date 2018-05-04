/* ResolveIP
 */

package main

import (
	"fmt"
	"net"
	"os"
)

// not include port
// ./bin/npwg-chapter3-3-ipaddr www.baidu.com
// ./bin/npwg-chapter3-3-ipaddr 31.13.86.16

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	// if the name is a hostname, it will at most return one IPAddr
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	// Resolved address is  31.13.86.16 ip 31.13.86.16
	fmt.Println("Resolved address is ", addr.String(), addr.Network(), addr.IP.String(), addr.Zone)
	os.Exit(0)
}
