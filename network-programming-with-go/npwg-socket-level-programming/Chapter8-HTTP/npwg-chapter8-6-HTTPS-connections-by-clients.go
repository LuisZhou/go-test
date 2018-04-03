package main

// HTTP + TLS
// Servers are required to return valid X.509 certificates before a client will accept data from them
// If the certificate is valid, then Go handles everything under the hood and the clients given previously run okay with
// https URLs.

// Many sites have invalid certificates. They may have expired, they may be self-signed, instead of by a recognised
// Certificate Authority or they may just have errors (such as having an incorrect server name)

// tr := &http.Transport{
// 	TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
// }
// client := &http.Client{Transport: transport}

func main() {

}
