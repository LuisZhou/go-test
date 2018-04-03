/* EchoServer
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	// "io"
	"golang.org/x/net/websocket"
)

// HTTP is a stream protocol. Web sockets are frame-based
// You prepare a block of data (of any size) and send it as a set of frames. Frames can contain either strings in UTF-8
// encoding or a sequence of bytes.

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing")

	for n := 0; n < 10; n++ {
		msg := "Hello  " + string(n+48)
		fmt.Println("Sending to client: " + msg)

		// The simplest way of using web sockets is just to prepare a block of data and ask the Go websocket library to
		// package it as a set of frame data, send them across the wire and receive it as the same block.

		// question: this may splite to multi frame?

		// you can also use the []byte{0, 1, 2}
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("Can't send")
			break
		}

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	// if you handle "/", means you handle everything.
	// if you handle others, only handle the matched url.
	// such as:
	// http.Handle("/test", websocket.Handler(Echo))
	// in client, you should specify the url in ws://localhost:12345/test
	// websocket.Dial("ws://localhost:12345/test", "", "http://localhost")
	// or you will get error:
	// Fatal error  websocket.Dial ws://localhost:12345: bad status
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
