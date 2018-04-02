/* ProxyGet
 */

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// URL requested should be the full URL of the destination
// HTTP header should contain a "Host" field, set to the proxy
// As long as the proxy is configured to pass such requests through, then that is all that needs to be done

// transport layer: http.Transport

// example
// ./bin/npwg-chapter8-5-1-Simple-proxy localhost:8087 http://www.bannedbook.com

// run deep into it!!!!
// http://colobu.com/2017/04/19/go-http-redirect/
// https://jonathanmh.com/tracing-preventing-http-redirects-golang/
// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
// If you don't have a suitable proxy to test this, then download and install the Squid proxy to your own computer.
// ref: https://haisum.github.io/2017/09/11/golang-ioutil-readall/
// http://colobu.com/2017/04/19/go-http-redirect/
// https://jonathanmh.com/tracing-preventing-http-redirects-golang/
// http://www.squid-cache.org/

// go can not:
// get proxy information from an  autoproxy.pac  file somewhere in your network.
// Linux systems using Gnome have a configuration system called __gconf__ in which proxy information can be stored: Go cannot access this.

// go only cat get env HTTP_PROXY or http_proxy
// func ProxyFromEnvironment(req *Request) (*url.URL, error)

// use httputil

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
		os.Exit(1)
	}

	proxyString := os.Args[1]
	proxyURL, err := url.Parse(proxyString)
	checkError(err)
	rawURL := os.Args[2]
	url, err := url.Parse(rawURL)
	checkError(err)

	// ProxyURL returns a proxy function (for use in a Transport)
	// that always returns the same URL.
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)} //question how about other shema?
	client := &http.Client{
		Transport: transport,
		// ref: https://gist.github.com/VojtechVitek/eb0171fc65f945a8641e
		CheckRedirect: func() func(req *http.Request, via []*http.Request) error {
			redirects := 0
			return func(req *http.Request, via []*http.Request) error {
				if redirects > 1000 {
					return errors.New("stopped after 1000 redirects")
				}
				redirects++
				return nil
			}
		}(),
	}

	request, err := http.NewRequest("GET", url.String(), nil)

	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	response, err := client.Do(request)

	checkError(err)
	fmt.Println("Read ok")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	fmt.Println("Reponse ok")

	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
