/* Head
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// summary:
// must has http://
// master ioutil.
// ref: https://golang.org/pkg/net/http/?m=all
// respond header is different with request header.

// __character set__: The server will deliver the content using some character set encoding, and possibly some transfer
// encoding. Usually this is a matter of __negotiation__ between user agent and server, but the simple  Get  command
// that we are using does not include the user agent component of the negotiation. So the server can send whatever
// character encoding it wishes.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	url := os.Args[1]

	// just get information about it
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}

	fmt.Println("-------------------\n")

	response2, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response2.Status != "200 OK" {
		fmt.Println(response2.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response2, false)
	fmt.Print(string(b))

	contentTypes := response2.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}

	body, err := ioutil.ReadAll(response2.Body)

	// question:
	// this will get EOF, why?

	// var buf [512]byte
	// reader := response2.Body
	// for {
	// 	n, err := reader.Read(buf[0:])
	// 	if err != nil {
	// 		fmt.Println("\nget error:", err)
	// 		os.Exit(0)
	// 	}
	// 	fmt.Print(string(buf[0:n]))
	// }

	fmt.Println(string(body), "\n")

	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	// each type is like [text/html; charset=UTF-8]
	// we want the UTF-8 only
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 || strings.Index(cType, "utf-8") != -1 {
			return true
		}
	}
	return false
}
