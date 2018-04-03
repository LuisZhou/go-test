/* EchoClient
 */
package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net"
	_ "net/http"
	"net/http/httptest"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	tlsServer := httptest.NewTLSServer(nil)
	tlsServerAddr := tlsServer.Listener.Addr().String()
	log.Print("Test TLS WebSocket server listening on ", tlsServerAddr)
	defer tlsServer.Close()
	// config, _ := NewConfig(fmt.Sprintf("wss://%s/echo", tlsServerAddr), "http://localhost")
	config, _ := websocket.NewConfig(service, "http://localhost")

	config.Dialer = &net.Dialer{
		Deadline: time.Now().Add(time.Minute * 30), // if use -time.Minute will get 'dial tcp: i/o timeout'
	}
	config.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := websocket.DialConfig(config)

	// config := tls.Config{RootCAs: nil, InsecureSkipVerify: true}
	// conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				break
			}
			fmt.Println("Couldn't receive msg " + err.Error())
			break
		}
		fmt.Println("Received from server: " + msg)
		// return the msg
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Coduln't return msg")
			break
		}
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
