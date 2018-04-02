package main

import (
	"fmt"
	"net/http"
	_ "net/url"
)

// a lower-level interface for user agents to communicate with HTTP servers.

// type Request

// https://golang.org/pkg/net/http/?m=all

func main() {
	request, err := http.NewRequest("GET", "http://baidu.com", nil)
	// the default set ISO-8859-1 always gets a value of one unless mentioned explicitly in the list
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	// return header will have: Content-Type: media type(such as text/html; charset=UTF-8),

	client := &http.Client{
	//CheckRedirect: redirectPolicyFunc,
	}

	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(request)
	fmt.Println(resp, err)
}
