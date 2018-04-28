/* PersonServerJSON
 */
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

// websocket clients and servers will exchange data in JSON format.
// Go object will be marshalled into JSON format.
// sent as a __UTF-8__ string.
// the receiver will read this string and unmarshal it back into a object.

type Person struct {
	Name   string
	Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
	var person Person
	err := websocket.JSON.Receive(ws, &person)
	if err != nil {
		fmt.Println("Can't receive")
	} else {
		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(ReceivePerson))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
