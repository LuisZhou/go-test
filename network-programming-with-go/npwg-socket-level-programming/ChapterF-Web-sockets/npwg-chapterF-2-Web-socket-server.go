package main

// When a request comes in that switches that connection to a being a web socket connection, the protocol handler is
// changed from an HTTP handler to a WebSocket handler

// it is only that TCP connection that gets its role changed: the server continues to be an HTTP server for other
// requests, while the TCP socket underlying that one connection is used as a web socket

func main() {
	// To handle web socket requests we simply register a different type of handler - a web socket handler.
	// Which handler the server uses is based on the URL pattern.

	http.Handle("/", websocket.Handler(WSHandler))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

// Question:
// How to make a web socket URL request?
