package main

// ## URLs specify the location of a resource

// resource:
// HTML document, an image, or a sound file, dynamically generated object

// really?
// When a user agent requests a resource, what is returned is not the resource itself, but some representation of that
// resource. For example, if the resource is a static file, then what is sent to the user agent is a __copy__ of the file.

// Multiple URLs may point to the same resource, and an HTTP server will return appropriate __representations__ of the
// resource for each URL.

// ## HTTP characteristics

// HTTP is a stateless, connectionless, reliable protocol
// In the simplest form, each request from a user agent is handled reliably and then the connection is broken.
// Each request involves a __separate__ TCP connection, so if many resources are required (such as images embedded in
// an HTML page) then many TCP connections have to be set up and torn down in a short space of time.

// many optimisations

// ## Versions

// Version 0.9 - totally obsolete
// Version 1.0 - almost obsolete
// Version 1.1 - current

// SP: SPACE?

// ## HTTP 0.9

// question
// Request = Simple-Request
// Simple-Request = "GET" SP Request-URI CRLF

// Response = Simple-Response
// Simple-Response = [Entity-Body]

// ## HTTP 1.0

// it was just left alongside the new version

// Request = Simple-Request | Full-Request
// Simple-Request = "GET" SP Request-URI CRLF
// Full-Request = Request-Line
//         *(General-Header
//         | Request-Header
//         | Entity-Header)
//         CRLF
//         [Entity-Body]

// Request-Line = Method SP Request-URI SP HTTP-Version CRLF
// Method = "GET" | "HEAD" | POST |
//     extension-method

// GET http://jan.newmarch.name/index.html HTTP/1.0

// Response = Simple-Response | Full-Response
// Simple-Response = [Entity-Body]
// Full-Response = Status-Line
//         *(General-Header
//         | Response-Header
//         | Entity-Header)
//         CRLF
//         [Entity-Body]

// Status-Line = HTTP-Version SP Status-Code SP Reason-Phrase CRLF

// Status-Code =      "200" ; OK
//         | "201" ; Created
//         | "202" ; Accepted
//         | "204" ; No Content
//         | "301" ; Moved permanently
//         | "302" ; Moved temporarily
//         | "304" ; Not modified
//         | "400" ; Bad request
//         | "401" ; Unauthorised
//         | "403" ; Forbidden
//         | "404" ; Not found
//         | "500" ; Internal server error
//         | "501" ; Not implemented
//         | "502" ; Bad gateway
//         | "503" | Service unavailable
//         | extension-code

// HTTP/1.0 200 OK

// The Entity-Header contains useful information about the Entity-Body to follow

// Entity-Header =    Allow
//         | Content-Encoding
//         | Content-Length
//         | Content-Type
//         | Expires
//         | Last-Modified
//         | extension-header

// HTTP/1.1 200 OK
// Date: Fri, 29 Aug 2003 00:59:56 GMT
// Server: Apache/2.0.40 (Unix)
// Accept-Ranges: bytes
// Content-Length: 1595
// Connection: close
// Content-Type: text/html; charset=ISO-8859-1

// ## HTTP 1.1

// fixes many problems with HTTP 1.0, but is more complex because of it.

// GET http://www.w3.org/index.html HTTP/1.1

// extending or refining the options
// there are more commands such as TRACE and CONNECT
// you should use absolute URLs, particularly for connecting by proxies
// there are more attributes such as If-Modified-Since, also for use by proxies

// changes:
// hostname identification (allows virtual hosts)
// content negotiation (multiple languages)
// persistent connections (reduces TCP overheads - this is very messy)
// chunked transfers
// byte ranges (request parts of documents)
// proxy support

// https://www.ietf.org/rfc/rfc1945.txt

func main() {

}
