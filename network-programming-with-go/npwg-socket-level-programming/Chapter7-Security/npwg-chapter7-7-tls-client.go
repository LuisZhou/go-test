/* TLSEchoClient
 */
package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

// you can not use ip 127.0.0.1 or 0.0.0.0 for your cer is not register that ip, if you use them, you get:
// Fatal error  x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs

// and you are not known authority, if you do not set InsecureSkipVerify to true in configure, you get:
// Fatal error  x509: certificate signed by unknown authority

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	// https://golang.org/pkg/crypto/tls/#Config

	// ref: https://groups.google.com/forum/#!topic/golang-nuts/v5ShM8R7Tdc
	// import ("net/http"; "crypto/tls")
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	// }
	// client := &http.Client{Transport: tr}
	// resp, err := client.Get("https://someurl:443/)

	conn, err := tls.Dial("tcp", service, &tls.Config{InsecureSkipVerify: true})
	checkError(err)

	for n := 0; n < 10; n++ {
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + string(n+48)))

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
