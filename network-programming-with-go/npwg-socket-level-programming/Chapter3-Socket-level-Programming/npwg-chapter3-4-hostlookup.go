package main

import (
	"fmt"
	"net"
	"os"
)

// example: ../../bin/npwg-chapter3-hostlookup www.baidu.com or
// provide it a host name, like ubuntu

// question: when lookup cname, it use port 53, why? -- DNS's port.

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	// this may return ipv6 too, if it exist.

	// todo: going on

	// func LookupAddr(addr string) (names []string, err error)
	// func LookupCNAME(host string) (cname string, err error)
	// func LookupHost(host string) (addrs []string, err error)
	// func LookupIP(host string) ([]IP, error)
	// func LookupMX(name string) ([]*MX, error)
	// func LookupNS(name string) ([]*NS, error)
	// func LookupPort(network, service string) (port int, err error)
	// func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
	// func LookupTXT(name string) ([]string, error)

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	// reference: https://en.wikipedia.org/wiki/CNAME_record

	// A Canonical Name record (abbreviated as CNAME record) is a type of resource record in the Domain Name System (DNS)
	// used to specify that a domain name is an alias for another domain (the 'canonical' domain).

	// This can prove convenient when running multiple services (like an FTP server and a webserver; each running on
	// different ports) from a single IP address. One can, for example, point ftp.example.com and www.example.com to the
	// DNS entry for example.com, which in turn has an A record which points to the IP address. Then, if the IP address
	// ever changes, one only has to record the change in one place within the network: in the DNS A record for example.com.

	// CNAME records must always point to another domain name, never directly to an IP address.

	// 	NAME                    TYPE   VALUE
	// --------------------------------------------------
	// bar.example.com.        CNAME  foo.example.com.
	// foo.example.com.        A      192.0.2.23

	cname, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("cname is ", cname)

	os.Exit(0)
}
