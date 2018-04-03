package main

// traditional model
// one request, one reply.

// backward
// + The first is that each request opens and closes a new TCP connection. HTTP 1.1 solved this by allowing persistent
//	 connections, so that a connection could be held open for a short period to allow for multiple requests (e.g. for
//	 images) to be made on the same server.
// + interaction model: event one form will be a requestion, displaying the response as a new page.
//	 __javascript__ can help in allowing error checking to be performed on form data before submission.
//	 __AJAX (Asynchronous JavaScript and XML)__ allows a browser to make a request and just use the response to update
//	 the display in place using the HTML Document Object Model (DOM).
// + AJAX just affects how the browser manages the returned pages.

// server to client
// only can do this with web socket: long-lived TCP connection to a Web sockets server.

// workflow
// + user agent sending a special HTTP request that says "switch to web sockets"
// + The TCP connection underlying the HTTP request is kept open, but both user agent and server switch to using the web
//	 sockets protocol instead of getting an HTTP response and closing the socket.

// ref:
// https://tools.ietf.org/html/draft-ietf-hybi-thewebsocketprotocol-17
// The client opens an HTTP connection and then replaces the HTTP protocol with its own WS protocol, re-using the same
// TCP connection.

// how
// how to do with persistent connections.
// how to switch from HTTP protocol to WS protocol.

func main() {

}
