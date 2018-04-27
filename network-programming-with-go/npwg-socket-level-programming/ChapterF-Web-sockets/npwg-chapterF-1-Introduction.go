package main

// r61 is miss.

// Traditional model
// one request, one reply.

// Backward
// + The first is that each request opens and closes a new TCP connection. HTTP 1.1 solved this by allowing persistent
//	 connections, so that a connection could be held open for a short period to allow for multiple requests (e.g. for
//	 images) to be made on the same server.

// + interaction model: event one form will be a requestion, displaying the response as a new page.
//	 __javascript__ can help in allowing error checking to be performed on form data before submission.
//	 __AJAX (Asynchronous JavaScript and XML)__ allows a browser to make a request and just use the response to update
//	 the display in place using the HTML Document Object Model (DOM).

// + AJAX just affects how the browser manages the returned pages.

// + What is missing is server initiated communications to the browser.

// Note
// There is no explicit extra support in Go for AJAX, as none is needed: the HTTP server just sees an ordinary HTTP POST
// request with possibly some XML or JSON data, and this.

// What the ws done.
// server to client
// long-lived TCP connection to a Web sockets server.

// workflow
// + user agent sending a special HTTP request that says "switch to web sockets": ws://XXX:XXX (this is a http request?)
// + The TCP connection underlying the HTTP request is kept open. (socket)
// + but both user agent and server switch to using the web sockets protocol instead of getting an HTTP response and closing the socket.

// ref:
// https://tools.ietf.org/html/draft-ietf-hybi-thewebsocketprotocol-17

// __how__
// how to do with http persistent connections.
// how to switch from HTTP protocol to WS protocol.

func main() {

}
