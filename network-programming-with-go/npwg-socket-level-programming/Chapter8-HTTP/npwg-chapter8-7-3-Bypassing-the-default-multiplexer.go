/* ServerHandler
 */

package main

import (
	"net/http"
)

func main() {
	// the multiplexer function
	myHandler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		// Just return no content - arbitrary headers can be set, arbitrary body
		rw.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":8080", myHandler)

	// Low-level servers
	// You first make a TCP server, and then wrap a  ServerConn  around it. Then you read  Request 's and write  Response 's
}
