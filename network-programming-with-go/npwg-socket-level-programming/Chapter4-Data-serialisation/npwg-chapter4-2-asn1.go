/* ASN.1
 */

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

// Abstract Syntax Notation One (ASN.1) was originally designed in 1984 for the telecommunications industry. ASN.1 is a
// complex standard, and a subset of it is supported by Go in the package "asn1". It builds self-describing serialised
// data from complex data structures. Its primary use in current networking systems is as the encoding for X.509
// certificates which are heavily used in authentication systems. The support in Go is based on what is needed to read
// and write X.509 certificates.

// two process, no matter what language, implement asn.1, we may consider them can communicate to each other, know what
// the package is meanning?

// func Marshal(val interface{}) ([]byte, os.Error)
// func Unmarshal(val interface{}, b []byte) (rest []byte, err os.Error)

type HelloWorld struct {
	Name string
}

func main() {
	// what type the asn1 consider about the '13'
	mdata, err := asn1.Marshal(13)
	checkError(err)

	// [2 1 13]
	fmt.Println(mdata)

	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", n)

	// support all the type?
	mdata2, err2 := asn1.Marshal(HelloWorld{
		Name: "what",
	})
	checkError(err2)

	// last four is what, how it descripe the type, is a table, and the first element is a string.
	// [48 6 19 4 119 104 97 116]
	fmt.Println(mdata2)

	// careful
	// this is application decide use what type to receive the unmarshal data.
	var h HelloWorld
	_, err3 := asn1.Unmarshal(mdata2, &h)
	checkError(err3)

	// After marshal/unmarshal:  {what1}
	// the name of the type is not cared by the asn1.
	fmt.Println("After marshal/unmarshal: ", h)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
