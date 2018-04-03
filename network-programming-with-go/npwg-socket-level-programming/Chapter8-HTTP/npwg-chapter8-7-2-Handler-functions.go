/* Print Env
 */

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// file handler for most files
	fileServer := http.FileServer(http.Dir("/var/www"))
	http.Handle("/", fileServer)

	// function handler for /cgi-bin/printenv
	http.HandleFunc("/cgi-bin/printenv", printEnv)

	// deliver requests to the handlers
	// The second argument to  HandleAndServe  could be  nil , and then calls are dispatched to all registered handlers
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
	// That's it!

	// A more specific pattern takes precedence over a more general pattern.

	// Go does have the ability to call external programs using  os.ForkExec , but does not yet have support for
	// dynamically linkable modules like Apache's  mod_perl.
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
