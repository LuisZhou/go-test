/* File Server
 */

package main

import (
	"fmt"
	"net/http"
	"os"
)

// Great !!!
// So easy.

// Go supplies a multi-plexer, that is, an object that will read and interpret requests. It hands out requests to
// handlers which run in their own thread. Thus much of the work of reading HTTP requests, decoding them and branching
// to suitable functions in their own thread is done for us.

func main() {
	// deliver files from the directory /var/www
	// fileServer := http.FileServer(http.Dir("/var/www"))
	// /home/httpd/html/
	fileServer := http.FileServer(http.Dir("/home/ubuntu/Documents/"))

	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8000", fileServer)
	checkError(err)
	// That's it!
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
